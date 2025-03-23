package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/michaelyusak/go-helper/apperror"
	"github.com/michaelyusak/kredit-plus-xyz/entity"
)

type userRepositoryPostgres struct {
	dbtx DBTX
}

func NewUserRepositoryPostgres(dbtx DBTX) *userRepositoryPostgres {
	return &userRepositoryPostgres{
		dbtx: dbtx,
	}
}

func (r *userRepositoryPostgres) RegisterUser(ctx context.Context, newUser entity.User) error {
	q := `
		INSERT INTO users (identity_number, full_name, legal_name, place_of_birth, date_of_birth, salary, identity_card_photo_url, selfie_photo_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $9);
	`

	now := time.Now().UnixMilli()

	_, err := r.dbtx.ExecContext(ctx, q,
		newUser.IdentityNumber,
		newUser.FullName,
		newUser.LegalName,
		newUser.PlaceOfBirth,
		newUser.DateOfBirth,
		newUser.Salary,
		newUser.IdentityCardPhotoUrl,
		newUser.SelfiePhotoUrl,
		now)
	if err != nil {
		return apperror.InternalServerError(fmt.Errorf("[UserRepository][RegisterUser][ExecContext] error: %w", err))
	}

	return nil
}

func (r *userRepositoryPostgres) GetOneByIdentityNumber(ctx context.Context, identityNumber string) (*entity.User, error) {
	q := `
		SELECT user_id, identity_number, full_name, legal_name, place_of_birth, date_of_birth, salary, identity_card_photo_url, selfie_photo_url, created_at, updated_at, deleted_at
		FROM users
		WHERE identity_number = $1;
	`

	var user entity.User

	err := r.dbtx.QueryRowContext(ctx, q, identityNumber).Scan(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, apperror.InternalServerError(fmt.Errorf("[UserRepository][GetOneByIdentityNumber][QueryRowContext] error: %w", err))
	}

	return &user, nil
}

func (r *userRepositoryPostgres) Lock(ctx context.Context) (error) {
	q := `
		LOCK TABLE users IN EXCLUSIVE MODE;
	`

	_, err := r.dbtx.ExecContext(ctx, q)
	if err != nil {
		return apperror.InternalServerError(fmt.Errorf("[UserRepository][Lock][ExecContext] error: %w", err))
	}

	return nil
}