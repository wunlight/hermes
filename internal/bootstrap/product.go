package bootstrap

import (
	"github.com/go-chi/chi/v5"
	"github.com/wunlight/hermes/internal/domains/product"
	"github.com/wunlight/hermes/internal/infrastructure/database"
	"github.com/wunlight/hermes/internal/middleware"
)

func Product(r chi.Router, db *database.DB, authMiddleware *middleware.AuthMiddleware) {
	repo := product.NewRepository(db.Executor())
	srv := product.NewService(repo)
	handler := product.NewHandler(srv)

	product.Mount(r, handler, authMiddleware)
}
