package bootstrap

import (
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/wunlight/hermes/internal/domains/auth"
	"github.com/wunlight/hermes/internal/domains/refresh_token"
	"github.com/wunlight/hermes/internal/domains/user"
	"github.com/wunlight/hermes/internal/infrastructure/database"
	"github.com/wunlight/hermes/internal/infrastructure/security/jwt"
	"github.com/wunlight/hermes/internal/infrastructure/security/password"
)

func Auth(r chi.Router, db *database.DB, tokenManager jwt.TokenManager) {
	userRepo := user.NewRepository(db.Executor())

	refreshTokenRepo := refresh_token.NewRepository(db.Executor())

	passwordHasher := password.NewArgon()

	authService := auth.NewService(userRepo, refreshTokenRepo, passwordHasher, tokenManager, 7*24*time.Hour)

	authHandler := auth.NewHandler(authService, "refresh_token", int((7 * 24 * time.Hour).Seconds()), false)

	auth.Mount(r, authHandler)
}
