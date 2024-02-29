package services

import (
	"context"

	"github.com/1BarCode/go-bank-v1/db"
)

type Services interface {
	CreateUser(ctx context.Context, username string) (db.User, error)
}

type services struct {
	// store db.Store
}

// func NewServices(store db.Store) Services {
// 	return &services{store: store}
// }
