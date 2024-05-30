package app

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	grpcapp "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/app/grpc"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/usecase"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/postgres"
)

const (
	_connTimeout = 5 * time.Second
)

type App struct {
	GRPCServer *grpcapp.App
}

func New() *App {

	// Config

	cfg := config.MustLoad()

	// Logger

	log := logger.NewLogger(cfg.Env)

	// Postgres

	ctx, cancel := context.WithTimeout(context.Background(), _connTimeout)
	defer cancel()
	pgxClient, err := postgres.NewPgxDatabase(ctx, cfg.CSAT.Postgres)
	if err != nil {
		panic("(csat) postgres connection error: " + err.Error())
	}

	// Layers

	csatRepo := repository.NewRepo(pgxClient)
	csatUsecase := usecase.NewUsecase(csatRepo)
	csatService := grpcapp.New(log, csatUsecase, cfg.CSAT)

	return &App{
		GRPCServer: csatService,
	}
}
