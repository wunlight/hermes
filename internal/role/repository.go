package role

import "context"

type Repository interface {
	Create(ctx context.Context, req CreateRequest) (*Role, error)
}
