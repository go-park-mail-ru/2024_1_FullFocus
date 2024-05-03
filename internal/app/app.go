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

	"github.com/gorilla/mux"

	authclient "github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/auth/grpc"
	profileclient "github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/profile/grpc"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	delivery "github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/http"
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
	config *config.Config
	server *server.Server
	router *mux.Router
	logger *slog.Logger
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

	// Server

	srv := server.NewServer(cfg.Main.Server, r)

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

	// csatConn, err := grpc.Dial(fmt.Sprintf("%s:%d", "localhost", 9090), grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	panic(err)
	// }

	// csatClient := gen.NewCSATClient(csatConn)
	// csatUsecase := usecase.NewCsatUsecase(csatClient)
	// csatHandler := delivery.NewCsatHandler(csatUsecase)
	// csatHandler.InitRouter(apiRouter)

	// Layers

	// Auth
	authUsecase := usecase.NewAuthUsecase(authClient, profileClient)
	authHandler := delivery.NewAuthHandler(authUsecase, cfg.SessionTTL)
	authHandler.InitRouter(apiRouter)

	// Cart
	cartRepo := repository.NewCartRepo(pgxClient)
	cartUsecase := usecase.NewCartUsecase(cartRepo)
	cartHandler := delivery.NewCartHandler(cartUsecase)
	cartHandler.InitRouter(apiRouter)

	// Profile
	profileUsecase := usecase.NewProfileUsecase(profileClient, cartRepo)
	profileHandler := delivery.NewProfileHandler(profileUsecase)
	profileHandler.InitRouter(apiRouter)

	// Order
	orderRepo := repository.NewOrderRepo(pgxClient)
	orderUsecase := usecase.NewOrderUsecase(orderRepo, cartRepo)
	orderHandler := delivery.NewOrderHandler(orderUsecase)
	orderHandler.InitRouter(apiRouter)

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
	reviewRepo := repository.NewReviewRepo(pgxClient)
	reviewUsecase := usecase.NewReviewUsecase(reviewRepo)
	reviewHandler := delivery.NewReviewHandler(reviewUsecase)
	reviewHandler.InitRouter(apiRouter)

	// Products
	productRepo := repository.NewProductRepo(pgxClient)
	productUsecase := usecase.NewProductUsecase(productRepo, categoryRepo)
	productHandler := delivery.NewProductHandler(productUsecase)
	productHandler.InitRouter(apiRouter)

	// Suggests
	suggestRepo := repository.NewSuggestRepo(elasticClient)
	suggestUsecase := usecase.NewSuggestUsecase(suggestRepo)
	suggestHandler := delivery.NewSuggestHandler(suggestUsecase)
	suggestHandler.InitRouter(apiRouter)

	// Middleware

	r.Use(middleware.NewLoggingMiddleware(log))
	r.Use(middleware.NewCORSMiddleware([]string{}))
	r.Use(middleware.NewAuthMiddleware(authClient))

	return &App{
		config: cfg,
		server: srv,
		router: r,
		logger: log,
	}
}

func (a *App) Run() {
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
