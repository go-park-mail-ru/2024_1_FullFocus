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
	"github.com/jackc/pgx/v5/pgxpool"
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

	dbCfg, err := pgxpool.ParseConfig(postgres.GetDSN(cfg.Postgres))
	if err != nil {
		panic("(review) postgres connection error: " + err.Error())
	}
	dbCfg.MaxConns = 10

	ctx, cancel := context.WithTimeout(context.Background(), _connTimeout)
	defer cancel()
	pgxClient, err := postgres.NewPgxDatabase(ctx, dbCfg)
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
