package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/michaelyusak/go-helper/apperror"
	"github.com/michaelyusak/go-helper/helper"
	"github.com/michaelyusak/kredit-plus-xyz/entity"
	"github.com/michaelyusak/kredit-plus-xyz/service"
)

type UserHandler struct {
	userService    service.UserService
	contextTimeout int64
}

func NewUserHandler(userService service.UserService, contextTimeout int64) *UserHandler {
	return &UserHandler{
		userService:    userService,
		contextTimeout: contextTimeout,
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var newUser entity.User

	data := ctx.Request.FormValue("data")
	if data == "" {
		ctx.Error(apperror.NewAppError(http.StatusBadRequest, errors.New("data is empty"), "data is empty"))
		return
	}

	err := json.Unmarshal([]byte(data), &newUser)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = validator.New().Struct(newUser)
	if err != nil {
		ctx.Error(err)
		return
	}

	identityCardPhotoFile, identityCardPhotoFileHeader, err := ctx.Request.FormFile("identity_card_photo")
	if err != nil {
		if identityCardPhotoFile != nil {
			ctx.Error(err)
			return
		}
	}

	selfiePhotoFile, selfiePhotoFileHeader, err := ctx.Request.FormFile("selfie_photo")
	if err != nil {
		if identityCardPhotoFile != nil {
			ctx.Error(err)
			return
		}
	}

	userMedia := entity.UserMedia{
		IdentityCardPhoto: entity.Media{
			File:   &identityCardPhotoFile,
			Header: identityCardPhotoFileHeader,
		},
		SelfiePhoto: entity.Media{
			File:   &selfiePhotoFile,
			Header: selfiePhotoFileHeader,
		},
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx.Request.Context(), time.Duration(h.contextTimeout)*time.Second)
	defer cancel()

	err = h.userService.Register(ctxWithTimeout, newUser, userMedia)
	if err != nil {
		ctx.Error(err)
		return
	}

	helper.ResponseOK(ctx, nil)
}
