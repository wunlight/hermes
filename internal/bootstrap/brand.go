package bootstrap

import (
	"github.com/go-chi/chi/v5"
	"github.com/wunlight/hermes/internal/domains/brand"
	"github.com/wunlight/hermes/internal/infrastructure/database"
	"github.com/wunlight/hermes/internal/middleware"
)

func Brand(r chi.Router, db *database.DB, authMiddleware *middleware.AuthMiddleware) {
	repo := brand.NewRepository(db.Executor())
	srv := brand.NewService(repo)
	handler := brand.NewHandler(srv)

	brand.Mount(r, handler, authMiddleware)
}
