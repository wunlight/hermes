package bootstrap

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/wunlight/hermes/internal/infrastructure/database"
	"github.com/wunlight/hermes/internal/infrastructure/security/jwt"
	"github.com/wunlight/hermes/internal/middleware"
)

func Register(r chi.Router, db *database.DB) {
	r.Get("/health", health)

	tokenManager := jwt.NewManager("your-secret", "hermes", 15*time.Minute)

	authMiddleware := middleware.NewAuthMiddleware(tokenManager)

	Auth(r, db, tokenManager)

	// Products Related
	Category(r, db, authMiddleware)
	Brand(r, db, authMiddleware)
	Unit(r, db, authMiddleware)
	Product(r, db, authMiddleware)

	Warehouse(r, db, authMiddleware)

	User(r, db, authMiddleware)
}
