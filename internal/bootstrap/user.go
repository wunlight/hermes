package bootstrap

import (
	"github.com/go-chi/chi/v5"
	"github.com/wunlight/hermes/internal/domains/user"
	"github.com/wunlight/hermes/internal/infrastructure/database"
	"github.com/wunlight/hermes/internal/middleware"
)

func User(r chi.Router, db *database.DB, authMiddleware *middleware.AuthMiddleware) {
	repo := user.NewRepository(db.Executor())

	service := user.NewService(repo)

	handler := user.NewHandler(service)

	user.Mount(r, handler, authMiddleware)
}
