package repositories

import (
	"contacts/models"
	"context"
	"gorm.io/gorm"
)

type userRepository struct {
	gorm *gorm.DB
}

func (u *userRepository) FindUserByEmail(context context.Context, email string) (models.User, error) {

	var user models.User

	err := u.gorm.
		WithContext(context).
		Raw(`SELECT * from users WHERE email = ?`, email).
		Scan(&user).
		Error

	return user, err
}

func (u *userRepository) CreateUser(context context.Context, user models.User) (models.User, error) {
	tx := u.gorm.WithContext(context).Create(&user)

	return user, tx.Error
}

func ProvideUserRepository(gorm *gorm.DB) UserRepository {

	return &userRepository{
		gorm: gorm,
	}
}
