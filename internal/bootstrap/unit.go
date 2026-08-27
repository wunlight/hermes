package bootstrap

import (
	"github.com/go-chi/chi/v5"
	"github.com/wunlight/hermes/internal/domains/unit"
	"github.com/wunlight/hermes/internal/infrastructure/database"
	"github.com/wunlight/hermes/internal/middleware"
)

func Unit(r chi.Router, db *database.DB, authMiddleware *middleware.AuthMiddleware) {
	repo := unit.NewRepository(db.Executor())
	srv := unit.NewService(repo)
	handler := unit.NewHandler(srv)

	unit.Mount(r, handler, authMiddleware)
}
