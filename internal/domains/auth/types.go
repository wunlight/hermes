package auth

import (
	"time"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Email    string
	Password string
	Name     string
}

type LoginRequest struct {
	Email    string
	Password string
}

type AuthResponse struct {
	UserID       uuid.UUID
	AccessToken  string
	RefreshToken string
}

type RegisterResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	UserID      uuid.UUID `json:"user_id"`
	AccessToken string    `json:"access_token"`
}

type RefreshResponse struct {
	UserID      uuid.UUID `json:"user_id"`
	AccessToken string    `json:"access_token"`
}
