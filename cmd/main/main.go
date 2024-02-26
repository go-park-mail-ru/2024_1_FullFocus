package main

import (
	"context"
	delivery "github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/http"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func run() {

	// Layers init

	userRepo := repository.NewUserRepo()
	sessionRepo := repository.NewSessionRepo()
	authUsecase := usecase.NewAuthUsecase(userRepo, sessionRepo)
	authHandler := delivery.NewAuthHandler(authUsecase)

	// Router init

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `Not found`, 404)
	})

	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.Handle("/login", http.HandlerFunc(authHandler.Login)).Methods("GET", "POST", "OPTIONS")
		auth.Handle("/signup", http.HandlerFunc(authHandler.Signup)).Methods("GET", "POST", "OPTIONS")
		auth.Handle("/logout", http.HandlerFunc(authHandler.Logout)).Methods("GET", "OPTIONS")
	}
	http.Handle("/", r)

	// Server init

	srv := http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	// Run server

	go func() {
		log.Printf("server is running...")
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("HTTP server ListenAndServe Error: %s", err.Error())
		}
	}()

	// Graceful shutdown

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	log.Printf("shutting down...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP Server Shutdown Error: %v", err)
	}
}

func main() {
	run()
}
