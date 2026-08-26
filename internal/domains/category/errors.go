package category

import "errors"

var (
	ErrNotFound              = errors.New("category not found")
	ErrCodeAlreadyExists     = errors.New("category code already exists")
	ErrInvalidParentID       = errors.New("invalid parent id")
	ErrInvalidID             = errors.New("invalid id")
	ErrCodeRequired          = errors.New("code is required")
	ErrNameRequired          = errors.New("name is required")
	ErrCategoryAlreadyExists = errors.New("category already exists")
)
