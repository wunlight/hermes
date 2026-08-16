package user

import "errors"

var (
	ErrNotFound           = errors.New("user not found")
	ErrInvalidID          = errors.New("invalid user id")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrEmailRequired      = errors.New("email is required")
	ErrInvalidEmail       = errors.New("invalid email")
	ErrNameRequired       = errors.New("name is required")
	ErrNameTooLong        = errors.New("name must not exceed 255 characters")
)
