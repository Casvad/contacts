package repositories

import (
	"contacts/dto"
	"contacts/models"
	"context"
)

type UserRepository interface {
	FindUserByEmail(context context.Context, email string) (models.User, error)
	CreateUser(context context.Context, user models.User) (models.User, error)
}

type UserContactRepository interface {
	CreateUserContact(context context.Context, contact models.UserContact) (models.UserContact, error)
	UpdateUserContact(context context.Context, contact models.UserContact) (models.UserContact, error)
	DeleteUserContact(context context.Context, contact models.UserContact) (models.UserContact, error)
	FindAllUserContactByUserId(context context.Context, userId int64) ([]dto.BasicUserContact, error)
	FindUserContactByUserIdAndId(context context.Context, userId, id int64) (models.UserContact, error)
}

type UserContactModificationRepository interface {
	CreateModification(context context.Context, contact models.UserContact) (models.UserContact, error)
}
