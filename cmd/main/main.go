package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	delivery "github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/http"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// @Title API Ozon
// @description Server API Ozon Application
// @version 1.0

// @host      http://62.233.46.235:8080
// @BasePath  /api

func run() {

	// Router init

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `Not found`, 404)
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	handler := c.Handler(r)

	// Server init

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           handler,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	// Layers init

	userRepo := repository.NewUserRepo()
	sessionRepo := repository.NewSessionRepo()
	authUsecase := usecase.NewAuthUsecase(userRepo, sessionRepo)
	authHandler := delivery.NewAuthHandler(srv, authUsecase)
	authHandler.InitRouter(r)

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
