package brand

import "errors"

var (
	ErrNotFound          = errors.New("brand not found")
	ErrCodeAlreadyExists = errors.New("brand code already exists")
	ErrInvalidID         = errors.New("invalid brand id")
	ErrCodeRequired      = errors.New("brand code required")
	ErrNameRequired      = errors.New("brand name required")
	ErrInvalidReq        = errors.New("invalid brand request")
)
