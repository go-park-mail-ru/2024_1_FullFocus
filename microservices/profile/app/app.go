package app

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	grpcapp "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/app/grpc"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/usecase"
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

	cfg.Logger.App = "profile"
	log := logger.NewLogger(cfg.Env, cfg.Logger)

	// Postgres

	ctx, cancel := context.WithTimeout(context.Background(), _connTimeout)
	defer cancel()
	pgxClient, err := postgres.NewPgxDatabase(ctx, cfg.Postgres)
	if err != nil {
		panic("(profile) postgres connection error: " + err.Error())
	}

	// Layers

	profileRepo := repository.NewRepo(pgxClient)
	profileUsecase := usecase.NewUsecase(profileRepo)
	profileService := grpcapp.New(log, profileUsecase, cfg.Profile)

	return &App{
		GRPCServer: profileService,
	}
}
