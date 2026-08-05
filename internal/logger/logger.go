package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	// if cfg.Env == "development" {
	return slog.New(slog.NewTextHandler(os.Stdout, opts))
	// }
	// return slog.New(slog.NewJSONHandler(os.Stdout, opts))
}
