package bootstrap

import (
	"github.com/go-chi/chi/v5"
	"github.com/wunlight/hermes/internal/domains/category"
	"github.com/wunlight/hermes/internal/infrastructure/database"
	"github.com/wunlight/hermes/internal/middleware"
)

func Category(r chi.Router, db *database.DB, authMiddleware *middleware.AuthMiddleware) {
	repo := category.NewRepository(db.Executor())
	srv := category.NewService(repo)
	handler := category.NewHandler(srv)

	category.Mount(r, handler, authMiddleware)
}
