package unit

import "errors"

var (
	ErrNotFound          = errors.New("unit not found")
	ErrCodeAlreadyExists = errors.New("unit code already exists")
	ErrInvalidID         = errors.New("invalid unit id")
	ErrCodeRequired      = errors.New("unit code required")
	ErrNameRequired      = errors.New("unit name required")
	ErrInvalidReq        = errors.New("invalid unit request")
)
