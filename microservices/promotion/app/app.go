package app

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	grpcapp "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/promotion/app/grpc"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/promotion/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/promotion/usecase"
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
	pgxClient, err := postgres.NewPgxDatabase(ctx, cfg.Postgres)
	if err != nil {
		panic("(promotion) postgres connection error: " + err.Error())
	}

	// Layers

	promotionRepo := repository.NewRepo(pgxClient)
	promotionUsecase := usecase.NewUsecase(promotionRepo)
	promotionService := grpcapp.New(log, promotionUsecase, cfg.Promotion)

	return &App{
		GRPCServer: promotionService,
	}
}
