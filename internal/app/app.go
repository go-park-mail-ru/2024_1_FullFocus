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

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	delivery "github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/http"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	corsmw "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/middleware/cors"
	logmw "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/middleware/logging"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/server"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
)

const (
	_timeout = 5 * time.Second
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

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `Not found`, 404)
	})

	// Middleware

	r.Use(logmw.NewLoggingMiddleware(log))
	r.Use(corsmw.NewCORSMiddleware([]string{}))

	// Server

	srv := server.NewServer(cfg.Server, r)

	// Layers

	// Auth
	userRepo := repository.NewUserRepo()
	sessionRepo := repository.NewSessionRepo(cfg.SessionTTL)
	authUsecase := usecase.NewAuthUsecase(userRepo, sessionRepo)
	authHandler := delivery.NewAuthHandler(authUsecase, cfg.SessionTTL)
	authHandler.InitRouter(apiRouter)

	// Products
	productRepo := repository.NewProductRepo()
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := delivery.NewProductHandler(productUsecase)
	productHandler.InitRouter(apiRouter)

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
			a.logger.Error(fmt.Sprintf("HTTP server ListenAndServe error: %s", err.Error()))
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
