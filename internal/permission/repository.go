package permission

import "context"

type Repository interface {
	Create(ctx context.Context, req CreateRequest) (*Permission, error)
	GetByCode(ctx context.Context, code string) (*Permission, error)
}
