package controllers

import (
	"contacts/dto"
	"contacts/models"
	"contacts/services"
	"github.com/gin-gonic/gin"
)

type userContactController struct {
	userContactService services.UserContactService
}

func (u *userContactController) GetById(context *gin.Context) {

	claims := getClaims(context)
	id := getId[int64](context)
	handleSimpleRequest(context, func() (models.UserContact, error) {
		return u.userContactService.FindByUserIdAndId(context, claims.UserId, id)
	})
}

func (u *userContactController) GetAllForUser(context *gin.Context) {

	claims := getClaims(context)
	handleSimpleRequest(context, func() ([]dto.BasicUserContact, error) {
		return u.userContactService.FindAllByUserId(context, claims.UserId)
	})
}

func (u *userContactController) Update(context *gin.Context) {
	claims := getClaims(context)
	id := getId[int64](context)
	handleRequestWithBody(context, func(request dto.UserContact) (models.UserContact, error) {
		return u.userContactService.Update(context, claims.UserId, id, request)
	})
}

func (u *userContactController) Delete(context *gin.Context) {

	claims := getClaims(context)
	id := getId[int64](context)
	handleSimpleRequest(context, func() (models.UserContact, error) {
		return u.userContactService.Delete(context, claims.UserId, id)
	})
}

func (u *userContactController) Create(context *gin.Context) {

	claims := getClaims(context)
	handleRequestWithBody(context, func(request dto.UserContact) (models.UserContact, error) {
		return u.userContactService.Create(context, claims.UserId, request)
	})
}

func ProvideUserContactController(userContactService services.UserContactService) UserContactController {

	return &userContactController{userContactService}
}
