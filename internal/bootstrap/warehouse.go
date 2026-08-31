package bootstrap

import (
	"github.com/go-chi/chi/v5"
	"github.com/wunlight/hermes/internal/domains/warehouse"
	"github.com/wunlight/hermes/internal/infrastructure/database"
	"github.com/wunlight/hermes/internal/middleware"
)

func Warehouse(r chi.Router, db *database.DB, authMiddleware *middleware.AuthMiddleware) {
	repo := warehouse.NewRepository(db.Executor())
	srv := warehouse.NewService(repo)
	handler := warehouse.NewHandler(srv)

	warehouse.Mount(r, handler, authMiddleware)
}
