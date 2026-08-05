package role

import (
	"context"

	"github.com/google/uuid"
	"github.com/wunlight/hermes/internal/permission"
)

type Repository interface {
	Create(ctx context.Context, req CreateRequest) (*Role, error)
	GetByCode(ctx context.Context, code string) (*Role, error)
	CreateRolePermission(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) error
}

type PermissionLookup interface {
	GetByCodes(
		ctx context.Context,
		codes []string,
	) ([]permission.Permission, error)
}
