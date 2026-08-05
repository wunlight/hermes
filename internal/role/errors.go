package role

import "errors"

var (
	ErrRoleCodeRequired  = errors.New("role code is required")
	ErrRoleNameRequired  = errors.New("role name is required")
	ErrRoleAlreadyExists = errors.New("role already exists")
)
