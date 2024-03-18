package logger

import (
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
