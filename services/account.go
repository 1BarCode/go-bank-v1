package services

import (
	"context"

	db "github.com/1BarCode/go-bank-v1/db/sqlc"
	"github.com/google/uuid"
)

func (s *services) CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error) {
	return s.store.CreateAccount(ctx, arg)
}

func (s *services) GetAccount(ctx context.Context, id uuid.UUID) (db.Account, error) {
	return s.store.GetAccount(ctx, id)
}

func (s *services) ListAccounts(ctx context.Context, arg db.ListAccountsParams) ([]db.Account, error) {
	return s.store.ListAccounts(ctx, arg)
}

func (s *services) DeleteAccount(ctx context.Context, id uuid.UUID) error {
	return s.store.DeleteAccount(ctx, id)
}