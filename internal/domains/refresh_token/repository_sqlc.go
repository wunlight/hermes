package refresh_token

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/wunlight/hermes/internal/infrastructure/adapters/pg_util"
	"github.com/wunlight/hermes/internal/infrastructure/adapters/sqlc"
)

type sqlcRepository struct {
	db sqlc.DBTX
}

func NewRepository(db sqlc.DBTX) Repository {
	return &sqlcRepository{db: db}
}

func (r *sqlcRepository) Create(
	ctx context.Context,
	token *RefreshToken,
) (*RefreshToken, error) {
	queries := sqlc.New(r.db)

	row, err := queries.CreateRefreshToken(
		ctx,
		sqlc.CreateRefreshTokenParams{
			UserID:    token.UserID,
			TokenHash: token.TokenHash,
			ExpiresAt: pg_util.TimeToTimestamptz(token.ExpiresAt),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("create refresh token: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) GetByHash(
	ctx context.Context,
	tokenHash string,
) (*RefreshToken, error) {
	queries := sqlc.New(r.db)

	row, err := queries.GetRefreshTokenByHash(ctx, tokenHash)
	if err != nil {
		return nil, fmt.Errorf("get refresh token by hash: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Revoke(
	ctx context.Context,
	id uuid.UUID,
) error {
	queries := sqlc.New(r.db)

	if err := queries.RevokeRefreshToken(ctx, id); err != nil {
		return fmt.Errorf("revoke refresh token: %w", err)
	}

	return nil
}

func toDomain(row sqlc.RefreshToken) *RefreshToken {
	return &RefreshToken{
		ID:        row.ID,
		UserID:    row.UserID,
		TokenHash: row.TokenHash,
		ExpiresAt: *pg_util.TimestamptzToTime(row.ExpiresAt),
		CreatedAt: *pg_util.TimestamptzToTime(row.CreatedAt),
		RevokedAt: pg_util.TimestamptzToTime(row.RevokedAt),
	}
}
