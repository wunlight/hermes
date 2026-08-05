package permission

import "errors"

var (
	ErrPermissionCodeRequired  = errors.New("permission code is required")
	ErrInvalidPermissionCode   = errors.New("invalid permission code format")
	ErrPermissionAlreadyExists = errors.New("permission already exists")
)
