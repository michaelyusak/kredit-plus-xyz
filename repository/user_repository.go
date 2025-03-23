package repository

import "database/sql"

type userRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) *userRepositoryPostgres {
	return &userRepositoryPostgres{
		db: db,
	}
}
