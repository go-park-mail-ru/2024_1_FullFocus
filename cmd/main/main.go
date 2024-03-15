package main

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	delivery "github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/http"
	corsmw "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/middleware/cors"
	logmw "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/middleware/logging"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
)

func run() {

	// Router init

	r := mux.NewRouter()
	apir := r.PathPrefix("/api").Subrouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `Not found`, 404)
	})

	// Logger

	l := logger.NewLogger()

	r.Use(corsmw.NewCORSMiddleware([]string{}))
	r.Use(logmw.NewLoggingMiddleware(l))

	// Server init

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	// Layers init

	// Auth
	userRepo := repository.NewUserRepo()
	sessionRepo := repository.NewSessionRepo()
	authUsecase := usecase.NewAuthUsecase(userRepo, sessionRepo)
	authHandler := delivery.NewAuthHandler(srv, authUsecase)
	authHandler.InitRouter(apir)

	// Products
	productRepo := repository.NewProductRepo()
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := delivery.NewProductHandler(srv, productUsecase)
	productHandler.InitRouter(apir)

	// Run server

	go func() {
		log.Printf("server is running...")
		if err := authHandler.Run(); err != nil {
			log.Printf("HTTP server ListenAndServe Error: %s", err.Error())
		}
	}()

	// Graceful shutdown

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	log.Printf("shutting down...")
	if err := authHandler.Stop(); err != nil {
		log.Printf("HTTP Server Shutdown Error: %v", err)
	}
}

func main() {
	run()
}
