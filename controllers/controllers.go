package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(context *gin.Context)
	Login(context *gin.Context)
}

type UserContactController interface {
	GetById(context *gin.Context)
	GetAllForUser(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	Create(context *gin.Context)
}
