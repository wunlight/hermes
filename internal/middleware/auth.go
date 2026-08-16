package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/wunlight/hermes/internal/infrastructure/security/jwt"
)

const userIDKey string = "user_id"

func UserID(ctx context.Context) (uuid.UUID, bool) {
	userID, ok := ctx.Value(userIDKey).(uuid.UUID)
	return userID, ok
}

type AuthMiddleware struct {
	tokenManager jwt.TokenManager
}

func NewAuthMiddleware(tokenManager jwt.TokenManager) *AuthMiddleware {
	return &AuthMiddleware{tokenManager: tokenManager}
}

func (m *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := extractBearerToken(r)
		if err != nil {
			unauthorized(w)
			return
		}

		claims, err := m.tokenManager.ParseAccessToken(token)
		if err != nil {
			unauthorized(w)
			return
		}

		userID, err := uuid.Parse(claims.Subject)
		if err != nil {
			unauthorized(w)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			userIDKey,
			userID,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}

func extractBearerToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")

	if header == "" {
		return "", errors.New(
			"authorization header is required",
		)
	}

	const scheme = "Bearer "

	if !strings.HasPrefix(header, scheme) {
		return "", errors.New(
			"authorization scheme must be Bearer",
		)
	}

	token := strings.TrimSpace(
		strings.TrimPrefix(header, scheme),
	)

	if token == "" {
		return "", errors.New(
			"access token is required",
		)
	}

	return token, nil
}

func unauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	_, _ = w.Write([]byte(`{"error":"unauthorized"}`))
}
