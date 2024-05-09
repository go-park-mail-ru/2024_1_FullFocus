package grpcapp

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/metrics"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	csatgrpc "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/grpc"
)

type App struct {
	logger     *slog.Logger
	gRPCServer *grpc.Server
	config     config.CSAT
	registry   *prometheus.Registry
}

func New(log *slog.Logger, csatService csatgrpc.CSAT, cfg config.CSAT) *App {
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
	reg := prometheus.NewRegistry()
	mt := metrics.NewMiddleware(metrics.NewMetrics(reg))
	gRPCServer := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			logger.StreamInterceptor(log),
		),
		grpc.ChainUnaryInterceptor(
			mt.UnaryInterceptor,
			logger.UnaryInterceptor(log),
			recovery.UnaryServerInterceptor(recoveryOpts...),
			logging.UnaryServerInterceptor(InterceptorLogger(log), loggingOpts...),
		),
	)
	csatgrpc.Register(gRPCServer, csatService)

	return &App{
		logger:     log,
		gRPCServer: gRPCServer,
		config:     cfg,
		registry:   reg,
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
		return fmt.Errorf("csat run error: %w", err)
	}
	a.logger.Info("grpc server started", slog.String("addr", l.Addr().String()))

	httpPort, _ := strconv.ParseInt(a.config.Server.Port, 10, 64)
	http.Handle("/metrics", promhttp.HandlerFor(a.registry, promhttp.HandlerOpts{Registry: a.registry}))
	go func() {
		if err = http.ListenAndServe(fmt.Sprintf(":%d", httpPort-1010), nil); err != nil {
			return
		}
	}()
	if err = a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("csat serve error: %w", err)
	}
	return nil
}

func (a *App) Stop() {
	a.logger.Info("stopping csat gRPC server...", slog.String("port", a.config.Server.Port))
	a.gRPCServer.GracefulStop()
	a.logger.Info("csat gRPC server stopped")
}
