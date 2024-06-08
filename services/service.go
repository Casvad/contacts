package services

import (
	"contacts/dto"
	"contacts/models"
	"context"
)

type UserService interface {
	Register(context context.Context, user dto.RegisterUser) (models.User, error)
	Login(context context.Context, login dto.Login) (dto.LoginResponse, error)
}

type UserContactService interface {
	Create(context context.Context, userId int64, userContact dto.UserContact) (models.UserContact, error)
	Update(context context.Context, userId, userContactId int64, userContact dto.UserContact) (models.UserContact, error)
	Delete(context context.Context, userId, userContactId int64) (models.UserContact, error)
	FindAllByUserId(context context.Context, userId int64) ([]dto.BasicUserContact, error)
	FindByUserIdAndId(context context.Context, userId, userContactId int64) (models.UserContact, error)
}
