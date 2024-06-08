package repositories

import (
	"contacts/dto"
	"contacts/models"
	"context"
	"gorm.io/gorm"
	"time"
)

type userContactRepository struct {
	gorm *gorm.DB
}

func (u *userContactRepository) CreateUserContact(context context.Context, contact models.UserContact) (models.UserContact, error) {

	err := u.gorm.WithContext(context).Create(&contact).Error

	return contact, err
}

func (u *userContactRepository) UpdateUserContact(context context.Context, contact models.UserContact) (models.UserContact, error) {
	err := u.gorm.WithContext(context).Model(contact).Select("*").Omit("created_at").Updates(contact).Error

	return contact, err

}

func (u *userContactRepository) DeleteUserContact(context context.Context, contact models.UserContact) (models.UserContact, error) {

	err := u.gorm.WithContext(context).Model(&contact).Update("deleted_at", time.Now()).Error

	return contact, err
}

func (u *userContactRepository) FindAllUserContactByUserId(context context.Context, userId int64) ([]dto.BasicUserContact, error) {

	var contacts []dto.BasicUserContact

	err := u.gorm.
		WithContext(context).
		Raw(`SELECT id, email, name from user_contacts WHERE user_id = ? and deleted_at is null`, userId).
		Scan(&contacts).
		Error

	if contacts == nil {
		contacts = []dto.BasicUserContact{}
	}

	return contacts, err
}

func (u *userContactRepository) FindUserContactByUserIdAndId(context context.Context, userId, id int64) (models.UserContact, error) {

	var userContact models.UserContact

	err := u.gorm.
		WithContext(context).
		Raw(`SELECT * from user_contacts WHERE user_id = ? and id = ?`, userId, id).
		Scan(&userContact).
		Error

	return userContact, err
}

func ProvideUserContactRepository(gorm *gorm.DB) UserContactRepository {

	return &userContactRepository{gorm}
}
