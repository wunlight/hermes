package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenManager interface {
	CreateAccessToken(userID uuid.UUID) (string, error)
	ParseAccessToken(token string) (*Claims, error)
}

type Claims struct {
	jwt.RegisteredClaims
}
