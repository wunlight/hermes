package jwt

import (
	"github.com/google/uuid"
)

type TokenManager interface {
	CreateAccessToken(userID uuid.UUID) (string, error)
	ParseAccessToken(token string) (*Claims, error)
}
