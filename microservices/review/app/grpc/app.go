package grpcapp

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	reviewgrpc "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/review/grpc"
)

type App struct {
	logger     *slog.Logger
	gRPCServer *grpc.Server
	config     config.Review
}

func New(log *slog.Logger, service reviewgrpc.Review, cfg config.Review) *App {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.StartCall, logging.FinishCall,
			logging.PayloadReceived, logging.PayloadSent,
		),
	}
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			log.Error("Recovered from panic", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}
	gRPCServer := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			logger.StreamInterceptor(log),
		),
		grpc.ChainUnaryInterceptor(
			logger.UnaryInterceptor(log),
			recovery.UnaryServerInterceptor(recoveryOpts...),
			logging.UnaryServerInterceptor(InterceptorLogger(log), loggingOpts...),
		),
	)
	reviewgrpc.Register(gRPCServer, service)

	return &App{
		logger:     log,
		gRPCServer: gRPCServer,
		config:     cfg,
	}
}

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", a.config.Server.Port))
	if err != nil {
		return fmt.Errorf("review run error: %w", err)
	}
	a.logger.Info("grpc server started", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("review serve error: %w", err)
	}
	return nil
}

func (a *App) Stop() {
	a.logger.Info("stopping review gRPC server...", slog.String("port", a.config.Server.Port))
	a.gRPCServer.GracefulStop()
	a.logger.Info("review gRPC server stopped")
}
