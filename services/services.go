package services

import (
	"context"
	"fmt"

	db "github.com/1BarCode/go-bank-v1/db/sqlc"
	"github.com/google/uuid"
)

type account interface {
	CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error)
	GetAccount(ctx context.Context, id uuid.UUID) (db.Account, error)
	ListAccounts(ctx context.Context, arg db.ListAccountsParams) ([]db.Account, error)
	DeleteAccount(ctx context.Context, id uuid.UUID) error
}

type transfer interface {
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

type user interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
}


type Services interface {
	account
	transfer
	user
}

type services struct {
	store db.Store
}

func NewServices(store db.Store) Services {
	return &services{store: store}
}

// execTx executes a function within a database transaction
func (s *services) execTx(ctx context.Context, cb func(*db.Queries) error) error {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx) // create new Queries object with transaction
	err = cb(q)   // execute the function with the new Queries object
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}