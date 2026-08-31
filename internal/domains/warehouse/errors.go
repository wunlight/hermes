package warehouse

import "errors"

var (
	ErrNotFound          = errors.New("warehouse not found")
	ErrCodeAlreadyExists = errors.New("warehouse code already exists")
	ErrInvalidID         = errors.New("invalid warehouse id")
	ErrCodeRequired      = errors.New("warehouse code required")
	ErrNameRequired      = errors.New("warehouse name required")
	ErrInvalidReq        = errors.New("invalid warehouse request")
)
