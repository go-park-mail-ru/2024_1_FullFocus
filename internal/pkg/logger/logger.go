package logger

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	file, _ := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{Level: slog.LevelDebug}))
}
