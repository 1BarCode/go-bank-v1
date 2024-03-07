package services

import (
	"context"

	db "github.com/1BarCode/go-bank-v1/db/sqlc"
	"github.com/google/uuid"
)

type account interface {
	CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error)
	GetAccount(ctx context.Context, id uuid.UUID) (db.Account, error)
	ListAccounts(ctx context.Context, arg db.ListAccountsParams) ([]db.Account, error)
	DeleteAccount(ctx context.Context, id uuid.UUID) error
}


type Services interface {
	account
}

type services struct {
	store db.Store
}

func NewServices(store db.Store) Services {
	return &services{store: store}
}
