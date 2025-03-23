package repository

import (
	"database/sql"
)

type Transaction interface {
	Begin() error
	Rollback() error
	Commit() error
	UserRepositoryPostgres() *userRepositoryPostgres
}

type sqlTransaction struct {
	db *sql.DB
	tx *sql.Tx
}

func NewSqlTransaction(db *sql.DB) *sqlTransaction {
	return &sqlTransaction{
		db: db,
	}
}

func (s *sqlTransaction) Begin() error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	s.tx = tx

	return nil
}

func (s *sqlTransaction) Rollback() error {
	return s.tx.Rollback()
}

func (s *sqlTransaction) Commit() error {
	return s.tx.Commit()
}

func (s *sqlTransaction) UserRepository() UserRepository {
	return &userRepositoryPostgres{
		dbtx: s.tx,
	}
}
