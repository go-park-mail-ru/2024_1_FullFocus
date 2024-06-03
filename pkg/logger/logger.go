package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

const (
	_envLocal = "local"
	_envProd  = "prod"
)

type ctxLogger struct{}

func NewLogger(env string, cfg config.Logger) *slog.Logger {
	log := &slog.Logger{}

	switch env {
	case _envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case _envProd:
		gelfWriter, err := gelf.NewTCPWriter(cfg.Target)
		if err != nil {
			panic(err)
		}
		log = slog.New(slog.NewJSONHandler(gelfWriter, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}
	return log.With("app_name", cfg.App)
}

func DefaultLogger() *slog.Logger {
	return NewLogger(_envProd, config.Logger{})
}

func Debug(ctx context.Context, msg string, args ...any) {
	getLoggerFromContext(ctx).Debug(msg, args...)
}

func Info(ctx context.Context, msg string, args ...any) {
	getLoggerFromContext(ctx).Info(msg, args...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	getLoggerFromContext(ctx).Warn(msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	getLoggerFromContext(ctx).Error(msg, args...)
}

// WithContext adds logger to context.
func WithContext(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// getLoggerFromContext returns logger from context.
func getLoggerFromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(ctxLogger{}).(*slog.Logger); ok {
		return l
	}
	return DefaultLogger()
}
