package helper

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
)

type ctxLogger struct{}

// ContextWithLogger adds logger to context.
func ContextWithLogger(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// GetLoggerFromContext returns logger from context.
func GetLoggerFromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(ctxLogger{}).(*slog.Logger); ok {
		return l
	}
	return logger.DefaultLogger()
}
