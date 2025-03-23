package service

import (
	"context"

	"github.com/michaelyusak/kredit-plus-xyz/entity"
)

type UserService interface {
	Register(ctx context.Context, newUser entity.User, newUserMedia entity.UserMedia) error
}

type TransactionService interface{}
