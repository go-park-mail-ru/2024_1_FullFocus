package app

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	grpcapp "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/auth/app/grpc"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/auth/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/auth/usecase"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/postgres"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/redis"
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
		panic("(auth) postgres connection error: " + err.Error())
	}

	// Redis

	redisClient := redis.NewClient(cfg.Auth.Redis)

	if err := redisClient.Ping().Err(); err != nil {
		panic("(auth) redis error: " + err.Error())
	}

	// Layers

	profileRepo := repository.NewAuthRepo(redisClient, pgxClient, cfg.SessionTTL)
	profileUsecase := usecase.NewAuthUsecase(profileRepo)
	authService := grpcapp.New(log, profileUsecase, cfg.Auth)

	return &App{
		GRPCServer: authService,
	}
}
