package services

import (
	"context"

	"github.com/1BarCode/go-bank-v1/db"
)

func (s *services) CreateUser(ctx context.Context, username string) (db.User, error) {
	user := db.User{
		Username: username,
	}
	// continue to call the store method here
	return user, nil
}