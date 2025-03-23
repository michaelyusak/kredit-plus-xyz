package repository

import (
	"context"
	"database/sql"

	"github.com/michaelyusak/kredit-plus-xyz/entity"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, newUser entity.User) error
	GetOneByIdentityNumber(ctx context.Context, identityNumber string) (*entity.User, error)
	Lock(ctx context.Context) (error)
}

type TransactionRepository interface{}

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
