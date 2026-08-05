package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"

	"github.com/wunlight/hermes/internal/config"
	"github.com/wunlight/hermes/internal/database"
	"github.com/wunlight/hermes/internal/logger"
)

func run() error {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	dbPool, err := database.New(ctx, cfg.DB)
	if err != nil {
		return fmt.Errorf("load database: %w", err)
	}

	_ = dbPool

	return nil
}

func main() {
	godotenv.Load()

	logger := logger.New()
	slog.SetDefault(logger)

	if err := run(); err != nil {
		slog.Error("application startup failed", "error", err)
		os.Exit(1)
	}
}
