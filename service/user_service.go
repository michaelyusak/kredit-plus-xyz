package service

import "github.com/michaelyusak/kredit-plus-xyz/users/repository"

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userServiceImpl {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}
