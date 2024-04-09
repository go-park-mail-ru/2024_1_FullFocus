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

	corsmw "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/middleware"
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	delivery "github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/http"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/server"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/minio"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/postgres"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/redis"
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

func Init() *App {
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

	// Middleware

	r.Use(corsmw.NewLoggingMiddleware(log))
	r.Use(corsmw.NewCORSMiddleware([]string{}))

	// Redis

	redisClient := redis.NewClient(cfg.Redis)

	if err := redisClient.Ping().Err(); err != nil {
		panic("redis error: " + err.Error())
	}

	// Minio

	minioClient, err := minio.NewClient(cfg.Minio)

	if err != nil {
		panic("minio connection error: " + err.Error())
	}

	// Postgres

	ctx, cancel := context.WithTimeout(context.Background(), _connTimeout)
	defer cancel()
	pgxClient, err := postgres.NewPgxDatabase(ctx, cfg.Postgres)
	if err != nil {
		panic("postgres connection error: " + err.Error())
	}

	// Server init

	srv := server.NewServer(cfg.Server, r)

	// Layers

	// Auth
	userRepo := repository.NewUserRepo(pgxClient)
	sessionRepo := repository.NewSessionRepo(redisClient, cfg.SessionTTL)
	authUsecase := usecase.NewAuthUsecase(userRepo, sessionRepo)
	authHandler := delivery.NewAuthHandler(authUsecase, cfg.SessionTTL)
	authHandler.InitRouter(apiRouter)

	// Product
	productRepo := repository.NewProductRepo(pgxClient)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := delivery.NewProductHandler(productUsecase)
	productHandler.InitRouter(apiRouter)

	// Order
	orderRepo := repository.NewOrderRepo(pgxClient)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)
	orderHandler := delivery.NewOrderHandler(orderUsecase)
	orderHandler.InitRouter(apiRouter)

	// Avatar
	avatarStorage := repository.NewAvatarStorage(minioClient)
	avatarUsecase := usecase.NewAvatarUsecase(avatarStorage, userRepo)
	avatarHandler := delivery.NewAvatarHandler(avatarUsecase)
	avatarHandler.InitRouter(apiRouter)

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
