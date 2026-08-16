package auth

import "errors"

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrRefreshTokenExpired = errors.New("refresh token expired")
	ErrRefreshTokenRevoked = errors.New("refresh token revoked")
	ErrEmailAlreadyExists  = errors.New("email already exists")
	ErrEmailRequired       = errors.New("email is required")
	ErrInvalidEmail        = errors.New("invalid email")
	ErrPasswordRequired    = errors.New("password is required")
	ErrPasswordTooShort    = errors.New("password must be at least 8 characters")
	ErrPasswordTooLong     = errors.New("password must not exceed 128 characters")
	ErrNameRequired        = errors.New("name is required")
	ErrNameTooLong         = errors.New("name must not exceed 255 characters")
)
