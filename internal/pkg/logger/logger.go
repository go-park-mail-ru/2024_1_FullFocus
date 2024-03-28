package logger

import (
	"context"
	"log/slog"
	"os"
)

const (
	_envLocal = "local"
	_envProd  = "prod"
)

const (
	_logFile = "logs/info.log"
)

type ctxLogger struct{}

func NewLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case _envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case _envProd:
		logFile, _ := os.OpenFile(_logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		log = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}
	return log
}

func DefaultLogger() *slog.Logger {
	return NewLogger(_envProd)
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
