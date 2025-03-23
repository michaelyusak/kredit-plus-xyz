package service

import (
	"context"
	"fmt"

	"github.com/michaelyusak/go-helper/apperror"
	"github.com/michaelyusak/kredit-plus-xyz/entity"
	"github.com/michaelyusak/kredit-plus-xyz/repository"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
	transaction repository.Transaction
}

func NewUserService(userRepository repository.UserRepository) *userServiceImpl {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

func (s *userServiceImpl) Register(ctx context.Context, newUser entity.User, newUserMedia entity.UserMedia) error {
	err := s.transaction.Begin()
	if err != nil {
		return apperror.InternalServerError(fmt.Errorf("[UserService][Register][transaction.Begin] error: %w", err))
	}

	userRepo := s.transaction.UserRepositoryPostgres()

	err = userRepo.Lock(ctx)
	if err != nil {
		return err
	}

	existing, err := userRepo.GetOneByIdentityNumber(ctx, newUser.IdentityNumber)
	if err != nil {
		return err
	}
	if existing != nil {
		return apperror.BadRequestError(fmt.Errorf("identity number registered"))
	}

	// save photo to storage
	// url, err := upload(newUserMedia)

	err = userRepo.RegisterUser(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}