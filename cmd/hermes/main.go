package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"github.com/wunlight/hermes/internal/config"
	"github.com/wunlight/hermes/internal/database"
	"github.com/wunlight/hermes/internal/logger"
	"github.com/wunlight/hermes/internal/router"
	"github.com/wunlight/hermes/internal/server"
)

func run() error {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	dbPool, err := database.New(ctx, cfg.DB)
	if err != nil {
		return fmt.Errorf("load database: %w", err)
	}
	defer dbPool.Close()

	r := router.New()
	router.Register(r, dbPool)

	srv := server.New(cfg.HTTP, r)
	serverErr := make(chan error, 1)

	go func() {
		slog.Info(
			"http server started",
			"addr", srv.Addr,
		)

		if err := srv.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
		close(serverErr)
	}()

	select {
	case err := <-serverErr:
		if err != nil {
			return fmt.Errorf("http server: %w", err)
		}
	case <-ctx.Done():
		slog.Info("shutdown signal received")
	}

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("shutdown server: %w", err)
	}

	slog.Info("http server stopped")

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
