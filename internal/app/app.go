package app

import (
	"context"
	"fmt"
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

func Run() {

	// Config

	cfg := config.MustLoad()

	// Logger

	log := logger.NewLogger(cfg.Env)

	// Router init

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `Not found`, 404)
	})

	// Middleware

	r.Use(logmw.NewLoggingMiddleware(log))
	r.Use(corsmw.NewCORSMiddleware([]string{}))

	// Server init

	srv := server.NewServer(cfg.Server, r)

	// Layers init

	// Auth
	userRepo := repository.NewUserRepo()
	sessionRepo := repository.NewSessionRepo()
	authUsecase := usecase.NewAuthUsecase(userRepo, sessionRepo)
	authHandler := delivery.NewAuthHandler(authUsecase)
	authHandler.InitRouter(apiRouter)

	// Products
	productRepo := repository.NewProductRepo()
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := delivery.NewProductHandler(productUsecase)
	productHandler.InitRouter(apiRouter)

	// Run server

	go func() {
		log.Info("server is running...")
		if err := srv.Run(); err != nil {
			log.Error(fmt.Sprintf("HTTP server ListenAndServe error: %s", err.Error()))
		}
	}()

	// Graceful shutdown

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	ctx, shutdown := context.WithTimeout(context.Background(), _timeout)
	defer shutdown()

	log.Info("shutting down...")
	if err := srv.Stop(ctx); err != nil {
		log.Error(fmt.Sprintf("HTTP server shutdown error: %v", err))
	}
}
