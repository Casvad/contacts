package repositories

import (
	"contacts/models"
	"contacts/utils/date"
	"context"
	"gorm.io/gorm"
	"time"
)

type userContactModificationRepository struct {
	gorm *gorm.DB
}

func (u *userContactModificationRepository) CreateModification(context context.Context, contact models.UserContact) (models.UserContact, error) {

	modification := models.UserContactModification{
		UserContactId: contact.Id,
		Modification:  models.UserContactMod(contact),
		CreatedAt:     date.ContactDate(time.Now()),
	}

	err := u.gorm.Create(&modification).Error

	return contact, err
}

func ProvideUserContactModificationRepository(
	db *gorm.DB,
) UserContactModificationRepository {

	return &userContactModificationRepository{db}
}
