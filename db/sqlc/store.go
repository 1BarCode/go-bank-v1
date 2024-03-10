package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

// Queries object used for single queries while db is used for transactions
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}

// execTx executes a function within a database transaction
// transaction is being moved to service layer
// TODO: need to remove this function and move the test to services layer
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx) // create new Queries object with transaction
	err = fn(q)   // execute the function with the new Queries object
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// expose Tx to services
func (store *SQLStore) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return store.db.BeginTx(ctx, opts)
}