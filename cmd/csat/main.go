package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/delivery"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/delivery/gen"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/usecase"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/postgres"
)

const (
	_connTimeout = 5 * time.Second
)

func run() {
	// Config

	cfg := config.MustLoad()

	// Logger

	log := logger.NewLogger(cfg.Env)

	// Postgres

	ctx, cancel := context.WithTimeout(context.Background(), _connTimeout)
	defer cancel()
	pgxClient, err := postgres.NewPgxDatabase(ctx, cfg.PostgresCSAT)
	if err != nil {
		panic("postgres connection error: " + err.Error())
	}

	// Layers

	csatRepo := repository.NewCSATRepo(pgxClient)
	csatUsecase := usecase.NewCSATUsecase(csatRepo)
	csatServer := delivery.NewCSATServer(csatUsecase)

	grpcServer := grpc.NewServer()
	gen.RegisterCSATServer(grpcServer, csatServer)

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 9090))
		if err != nil {
			log.Error("listen error: " + err.Error())
		}
		log.Info("grpc server started", slog.String("addr", listener.Addr().String()))
		if err = grpcServer.Serve(listener); err != nil {
			log.Error("serve returned err: " + err.Error())
		}
	}()

	// Graceful shutdown

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	log.Info("shutting down...")
	grpcServer.GracefulStop()
}

func main() {
	run()
}
