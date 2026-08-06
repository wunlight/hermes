package database

import (
	"context"
	"fmt"

	"github.com/wunlight/hermes/internal/adapters/sqlc"
)

type Executor = sqlc.DBTX

type Manager interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context, tx Executor) error) error
}

func (db *DB) WithinTransaction(ctx context.Context, fn func(ctx context.Context, tx Executor) error) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	if err := fn(ctx, tx); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	committed = true

	return nil
}
