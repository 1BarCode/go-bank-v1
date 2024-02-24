package db

import (
	"context"
)

type CreateUserParams struct {
	Username string
}

type UserQuerier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
}