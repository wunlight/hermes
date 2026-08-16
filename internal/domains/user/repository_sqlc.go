package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/wunlight/hermes/internal/infrastructure/adapters/pg_util"
	"github.com/wunlight/hermes/internal/infrastructure/adapters/sqlc"
)

type sqlcRepository struct {
	db sqlc.DBTX
}

func NewRepository(db sqlc.DBTX) Repository {
	return &sqlcRepository{
		db: db,
	}
}

func (r *sqlcRepository) Create(ctx context.Context, user *User) (*User, error) {
	queries := sqlc.New(r.db)

	row, err := queries.CreateUser(
		ctx,
		sqlc.CreateUserParams{
			Email:        user.Email,
			PasswordHash: user.PasswordHash,
			Name:         user.Name,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) GetByID(ctx context.Context, id uuid.UUID) (*User, error) {
	queries := sqlc.New(r.db)

	row, err := queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	queries := sqlc.New(r.db)

	row, err := queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("get user by email: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Update(ctx context.Context, user *User) (*User, error) {
	queries := sqlc.New(r.db)

	row, err := queries.UpdateUser(
		ctx,
		sqlc.UpdateUserParams{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("update user: %w", err)
	}

	return toDomain(row), nil
}

func (r *sqlcRepository) Delete(ctx context.Context, id uuid.UUID) error {
	queries := sqlc.New(r.db)

	if err := queries.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}

func (r *sqlcRepository) List(ctx context.Context) ([]*User, error) {
	queries := sqlc.New(r.db)

	rows, err := queries.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}

	users := make([]*User, 0, len(rows))

	for _, row := range rows {
		users = append(users, toDomain(row))
	}

	return users, nil
}

func toDomain(row sqlc.User) *User {
	return &User{
		ID:           row.ID,
		Email:        row.Email,
		PasswordHash: row.PasswordHash,
		Name:         row.Name,
		CreatedAt:    *pg_util.TimestamptzToTime(row.CreatedAt),
		UpdatedAt:    *pg_util.TimestamptzToTime(row.UpdatedAt),
		DeletedAt:    pg_util.TimestamptzToTime(row.DeletedAt),
	}
}
