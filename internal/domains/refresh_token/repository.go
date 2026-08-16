package refresh_token

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, token *RefreshToken) (*RefreshToken, error)
	GetByHash(ctx context.Context, tokenHash string) (*RefreshToken, error)
	Revoke(ctx context.Context, id uuid.UUID) error
}
