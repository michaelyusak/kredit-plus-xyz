package repository

import "database/sql"

type transactionRepositoryPostgres struct {
	db *sql.DB
}

func NewTransactionRepositoryPostgres(db *sql.DB) *transactionRepositoryPostgres {
	return &transactionRepositoryPostgres{
		db: db,
	}
}