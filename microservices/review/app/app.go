package app

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	grpcapp "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/review/app/grpc"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/review/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/review/usecase"
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

	cfg.Logger.App = "review"
	log := logger.NewLogger(cfg.Env, cfg.Logger)

	// Postgres

	ctx, cancel := context.WithTimeout(context.Background(), _connTimeout)
	defer cancel()
	pgxClient, err := postgres.NewPgxDatabase(ctx, cfg.Postgres)
	if err != nil {
		panic("(review) postgres connection error: " + err.Error())
	}

	// Layers

	reviewRepo := repository.NewRepo(pgxClient)
	reviewUsecase := usecase.NewUsecase(reviewRepo)
	profileService := grpcapp.New(log, reviewUsecase, cfg.Review)

	return &App{
		GRPCServer: profileService,
	}
}
