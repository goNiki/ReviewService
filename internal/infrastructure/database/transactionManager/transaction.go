package txmanager

import (
	"context"

	"github.com/goNiki/ReviewService/internal/infrastructure/database"
)

type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context, tx database.Tx) error) error
}

type TxManager struct {
	db *database.Db
}

func NewTransactionManager(db *database.Db) *TxManager {
	return &TxManager{db: db}
}

func (tm *TxManager) WithTransaction(
	ctx context.Context,
	fn func(ctx context.Context, tx database.Tx) error,
) error {
	tx, err := tm.db.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := fn(ctx, tx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
