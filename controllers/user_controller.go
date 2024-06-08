package controllers

import (
	"contacts/dto"
	"contacts/models"
	"contacts/services"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userService services.UserService
}

func (u *userController) Register(context *gin.Context) {

	handleRequestWithBody(context, func(request dto.RegisterUser) (models.User, error) {
		return u.userService.Register(context, request)
	})
}

func (u *userController) Login(context *gin.Context) {

	handleRequestWithBody(context, func(request dto.Login) (dto.LoginResponse, error) {
		return u.userService.Login(context, request)
	})
}

func ProvideUserController(userService services.UserService) UserController {

	return &userController{userService}
}
