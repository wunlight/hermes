package bootstrap

import (
	"github.com/go-chi/chi/v5"
	"github.com/wunlight/hermes/internal/database"
)

func Register(r chi.Router, db *database.DB) {
	r.Get("/health", health)

	Permission(r, db)
}
