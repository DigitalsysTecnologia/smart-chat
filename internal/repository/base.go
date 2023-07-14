package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

const DbTransactionKey string = "transaction_key"

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{
		db: db,
	}
}

func (b *BaseRepository) storeTransaction(ctx context.Context, connWithTransaction *gorm.DB) (context.Context, error) {
	newCtx := context.WithValue(ctx, DbTransactionKey, connWithTransaction)
	return newCtx, nil
}

func (b *BaseRepository) getTransaction(ctx context.Context) (*gorm.DB, error) {
	connWithTransaction := ctx.Value(DbTransactionKey)
	if connWithTransaction == nil {
		return nil, nil
	}
	return connWithTransaction.(*gorm.DB), nil
}

// BeginTransaction - Generates a new DB transaction and insert it on the context.
func (b *BaseRepository) BeginTransaction(ctx context.Context) (context.Context, error) {
	connWithTransaction := b.db.Begin()
	return b.storeTransaction(ctx, connWithTransaction)
}

func (b *BaseRepository) RollbackTransaction(ctx context.Context) error {
	connWithTransaction, err := b.getTransaction(ctx)
	if err != nil {
		return err
	}

	if connWithTransaction == nil {
		return errors.New("no transaction found on current context")
	}
	return connWithTransaction.Rollback().Error
}

func (b *BaseRepository) CommitTransaction(ctx context.Context) error {
	connWithTransaction, err := b.getTransaction(ctx)
	if err != nil {
		return err
	}

	if connWithTransaction == nil {
		return errors.New("no transaction found on current context")
	}
	return connWithTransaction.Commit().Error
}

func (b *BaseRepository) GetConnection(ctx context.Context) (*gorm.DB, error) {
	var result *gorm.DB
	connWithTransaction, err := b.getTransaction(ctx)

	if err != nil {
		return nil, err
	}

	if connWithTransaction != nil {
		result = connWithTransaction
	} else {
		result = b.db
	}

	return result.WithContext(ctx), nil
}
