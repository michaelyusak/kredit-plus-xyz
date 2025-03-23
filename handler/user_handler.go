package handler

import "github.com/michaelyusak/kredit-plus-xyz/users/service"

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}