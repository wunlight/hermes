package role

import "errors"

var (
	ErrRoleCodeRequired         = errors.New("role code is required")
	ErrRoleNameRequired         = errors.New("role name is required")
	ErrRoleAlreadyExists        = errors.New("role already exists")
	ErrPermissionNotFound       = errors.New("permission not found")
	ErrDuplicatePermissionCodes = errors.New("duplicate permission codes")
	ErrPermissionCodeRequired   = errors.New("permission code is required")
)
