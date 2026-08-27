package unit

import "errors"

var (
	ErrNotFound          = errors.New("product not found")
	ErrCodeAlreadyExists = errors.New("product code already exists")
	ErrInvalidID         = errors.New("invalid product id")
	ErrCodeRequired      = errors.New("product code required")
	ErrNameRequired      = errors.New("product name required")
	ErrInvalidReq        = errors.New("invalid product request")
)
