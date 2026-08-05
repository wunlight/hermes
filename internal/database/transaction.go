package database

import (
	"context"

	"github.com/wunlight/hermes/internal/adapters/sqlc"
)

type TxFunc func(ctx context.Context, db sqlc.DBTX) error

func (db *DB) Transaction(ctx context.Context, fn TxFunc) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := fn(ctx, tx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
