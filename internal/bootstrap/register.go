package bootstrap

import (
	"github.com/go-chi/chi/v5"
	"github.com/wunlight/hermes/internal/infrastructure/database"
)

func Register(r chi.Router, db *database.DB) {
	r.Get("/health", health)

	Auth(r, db)
}
