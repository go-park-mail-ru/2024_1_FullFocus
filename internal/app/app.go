package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/centrifuge"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/metrics"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	authclient "github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/auth/grpc"
	csatclient "github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/csat/grpc"
	profileclient "github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/profile/grpc"
	promotionclient "github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/promotion/grpc"
	reviewclient "github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/review/grpc"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	delivery "github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/http"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/cache"
	elasticsetup "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/elasticsearch"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/middleware"
	miniosetup "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/minio"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/server"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/elasticsearch"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/minio"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/postgres"
)

const (
	_timeout     = 5 * time.Second
	_connTimeout = 10 * time.Second
)

type App struct {
	config   *config.Config
	server   *server.Server
	router   *mux.Router
	logger   *slog.Logger
	registry *prometheus.Registry
}

func MustInit() *App {
	// Config

	cfg := config.MustLoad()

	// Logger

	log := logger.NewLogger(cfg.Env)

	// Router

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, `Not found`, 404)
	})

	// Minio

	minioClient, err := minio.NewClient(cfg.Minio)

	if err != nil {
		panic("minio connection error: " + err.Error())
	}

	if err = miniosetup.InitBucket(context.Background(), minioClient, cfg.Minio.AvatarBucket); err != nil {
		panic("minio init bucket error: " + err.Error())
	}

	// Postgres

	ctx, cancel := context.WithTimeout(context.Background(), _connTimeout)
	defer cancel()

	pgxClient, err := postgres.NewPgxDatabase(ctx, cfg.Postgres)
	if err != nil {
		panic("postgres connection error: " + err.Error())
	}

	// Elasticsearch

	elasticClient, err := elasticsearch.NewClient(cfg.Elasticsearch)
	if err != nil {
		panic("elasticsearch connection error: " + err.Error())
	}

	_, err = elasticClient.Ping()
	if err != nil {
		panic("elasticsearch ping error: " + err.Error())
	}

	ctx, cancel = context.WithTimeout(context.Background(), _timeout)
	defer cancel()

	if err = elasticsetup.InitElasticData(ctx, pgxClient, elasticClient); err != nil {
		panic("elasticsearch init data error: " + err.Error())
	}

	// Centrifugo

	ctx, cancel = context.WithTimeout(context.Background(), _timeout)
	defer cancel()

	centrifugoClient := centrifuge.NewCentrifugeClient(ctx, cfg.Centrifugo)
	if centrifugoClient == nil {
		panic("centrifugo connection error")
	}

	// Server

	srv := server.NewServer(cfg.Main.Server, r)

	// Cache

	promotionCache := cache.NewPromoProductsCache()

	// Clients

	// Auth
	ctx, cancel = context.WithTimeout(context.Background(), _connTimeout)
	defer cancel()

	authClient, err := authclient.New(ctx, log, cfg.Main.Clients.AuthClient)
	if err != nil {
		panic("auth service connection error: " + err.Error())
	}

	// Profile
	ctx, cancel = context.WithTimeout(context.Background(), _connTimeout)
	defer cancel()

	profileClient, err := profileclient.New(ctx, log, cfg.Main.Clients.ProfileClient)
	if err != nil {
		panic("profile service connection error: " + err.Error())
	}

	// CSAT
	ctx, cancel = context.WithTimeout(context.Background(), _connTimeout)
	defer cancel()

	csatClient, err := csatclient.New(ctx, log, cfg.Main.Clients.CSATClient)
	if err != nil {
		panic("csat service connection error: " + err.Error())
	}

	// Review
	ctx, cancel = context.WithTimeout(context.Background(), _connTimeout)
	defer cancel()

	reviewClient, err := reviewclient.New(ctx, log, cfg.Main.Clients.ReviewClient)
	if err != nil {
		panic("review service connection error: " + err.Error())
	}

	// Promotion
	ctx, cancel = context.WithTimeout(context.Background(), _connTimeout)
	defer cancel()

	promotionClient, err := promotionclient.New(ctx, log, cfg.Main.Clients.PromotionClient)
	if err != nil {
		panic("promotion service connection error: " + err.Error())
	}

	// Layers

	// Auth
	authUsecase := usecase.NewAuthUsecase(authClient, profileClient)
	authHandler := delivery.NewAuthHandler(authUsecase, cfg.SessionTTL)
	authHandler.InitRouter(apiRouter)

	// Cart
	cartRepo := repository.NewCartRepo(pgxClient)
	cartUsecase := usecase.NewCartUsecase(cartRepo, promotionClient)
	cartHandler := delivery.NewCartHandler(cartUsecase)
	cartHandler.InitRouter(apiRouter)

	// Promocode
	promocodeRepo := repository.NewPromocodeRepo(pgxClient)
	promocodeUsecase := usecase.NewPromocodeUsecase(promocodeRepo)
	promocodeHandler := delivery.NewPromocodeHandler(promocodeUsecase)
	promocodeHandler.InitRouter(apiRouter)

	// Avatar
	avatarStorage := repository.NewAvatarStorage(minioClient, cfg.Minio)
	avatarUsecase := usecase.NewAvatarUsecase(avatarStorage, profileClient)
	avatarHandler := delivery.NewAvatarHandler(avatarUsecase)
	avatarHandler.InitRouter(apiRouter)

	// Categories
	categoryRepo := repository.NewCategoryRepo(pgxClient)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryHandler := delivery.NewCategoryHandler(categoryUsecase)
	categoryHandler.InitRouter(apiRouter)

	// Reviews
	reviewUsecase := usecase.NewReviewUsecase(profileClient, reviewClient, promotionCache)
	reviewHandler := delivery.NewReviewHandler(reviewUsecase)
	reviewHandler.InitRouter(apiRouter)

	// Products
	productRepo := repository.NewProductRepo(pgxClient)
	productUsecase := usecase.NewProductUsecase(productRepo, categoryRepo, promotionClient)
	productHandler := delivery.NewProductHandler(productUsecase)
	productHandler.InitRouter(apiRouter)

	// Notifications
	notificationRepo := repository.NewNotificationRepo(pgxClient)
	notificationUsecase := usecase.NewNotificationUsecase(notificationRepo, centrifugoClient)
	notificationHandler := delivery.NewNotificationHandler(notificationUsecase)
	notificationHandler.InitRouter(apiRouter)

	// Profile
	profileUsecase := usecase.NewProfileUsecase(authClient, profileClient, cartRepo, promocodeRepo, notificationRepo)
	profileHandler := delivery.NewProfileHandler(profileUsecase)
	profileHandler.InitRouter(apiRouter)

	// Order
	orderRepo := repository.NewOrderRepo(pgxClient)
	orderUsecase := usecase.NewOrderUsecase(orderRepo, cartRepo, productRepo, promocodeRepo, notificationRepo, promotionClient)
	orderHandler := delivery.NewOrderHandler(orderUsecase, notificationUsecase)
	orderHandler.InitRouter(apiRouter)

	// Suggests
	suggestRepo := repository.NewSuggestRepo(elasticClient)
	suggestUsecase := usecase.NewSuggestUsecase(suggestRepo)
	suggestHandler := delivery.NewSuggestHandler(suggestUsecase)
	suggestHandler.InitRouter(apiRouter)

	// CSAT
	csatUsecase := usecase.NewCsatUsecase(csatClient)
	csatHandler := delivery.NewCsatHandler(csatUsecase)
	csatHandler.InitRouter(apiRouter)

	// Promotion
	promotionUsecase := usecase.NewPromotionUsecase(ctx, productRepo, promotionClient, promotionCache)
	promotionHandler := delivery.NewPromotionHandler(promotionUsecase)
	promotionHandler.InitRouter(apiRouter)

	// Middleware
	reg := prometheus.NewRegistry()
	r.Use(middleware.NewLoggingMiddleware(metrics.NewMetrics(reg), log))
	r.Use(middleware.NewCORSMiddleware([]string{}))
	r.Use(middleware.NewAuthMiddleware(authClient))
	r.Use(middleware.NewAuthorizationMiddleware(cfg.AccessToken))

	return &App{
		config:   cfg,
		server:   srv,
		router:   r,
		logger:   log,
		registry: reg,
	}
}

func (a *App) Run() {
	a.router.Handle("/public/metrics", promhttp.HandlerFor(
		a.registry,
		promhttp.HandlerOpts{
			Registry: a.registry,
		},
	))

	go func() {
		a.logger.Info("server is running...")
		if err := a.server.Run(); err != nil {
			a.logger.Error("HTTP server ListenAndServe error: " + err.Error())
		}
	}()

	// Graceful shutdown

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	ctx, shutdown := context.WithTimeout(context.Background(), _timeout)
	defer shutdown()

	a.logger.Info("shutting down...")
	if err := a.server.Stop(ctx); err != nil {
		a.logger.Error(fmt.Sprintf("HTTP server shutdown error: %v", err))
	}
}
