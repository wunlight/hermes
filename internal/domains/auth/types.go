package auth

import "github.com/google/uuid"

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
