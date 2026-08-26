package brand

import "errors"

var (
	ErrNotFound          = errors.New("brand not found")
	ErrCodeAlreadyExists = errors.New("brand code already exists")
)
