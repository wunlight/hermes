package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	return slog.New(slog.NewTextHandler(os.Stdout, opts))
}
