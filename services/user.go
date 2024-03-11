package services

import (
	"context"

	db "github.com/1BarCode/go-bank-v1/db/sqlc"
)

func (s *services) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return s.store.CreateUser(ctx, arg)
}