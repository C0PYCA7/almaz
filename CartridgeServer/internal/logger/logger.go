package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	if os.Getenv("ENV") != "local" {
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
}
