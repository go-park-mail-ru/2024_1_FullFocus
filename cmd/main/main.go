package main

import (
	middleware "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/middleware/cors"
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
)

func run() {

	// Router init

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `Not found`, 404)
	})

	r.Use(middleware.CORSMiddleware([]string{}))

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
