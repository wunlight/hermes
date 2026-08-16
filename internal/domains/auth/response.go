package auth

import (
	"time"

	"github.com/google/uuid"
)

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
