package services

import (
	"contacts/dto"
	"contacts/models"
	"contacts/repositories"
	"contacts/utils/date"
	"contacts/utils/errors"
	"context"
	"fmt"
	"net/http"
	"time"
)

type userContactService struct {
	userContactRepository             repositories.UserContactRepository
	userContactModificationRepository repositories.UserContactModificationRepository
}

func (u *userContactService) Create(ctx context.Context, userId int64, userContact dto.UserContact) (models.UserContact, error) {

	userContactDB, err := u.userContactRepository.CreateUserContact(ctx, models.UserContact{
		UserId:      userId,
		Email:       userContact.Email,
		Name:        userContact.Name,
		PrefixPhone: userContact.PrefixPhone,
		Phone:       userContact.Phone,
		CreatedAt:   date.ContactDate(time.Now()),
		UpdatedAt:   date.ContactDate(time.Now()),
	})

	if err != nil {
		return userContactDB, errors.NewUserErrorWithError("user_contact.error", http.StatusBadRequest, err)
	}

	go func() {
		_, err := u.userContactModificationRepository.CreateModification(context.Background(), userContactDB)
		if err != nil {
			fmt.Printf("Error while creating modification on create %v for user: %d", err, userId)
		}
	}()

	return userContactDB, nil
}

func (u *userContactService) Update(ctx context.Context, userId, userContactId int64, userContact dto.UserContact) (models.UserContact, error) {

	userContactDB, err := u.userContactRepository.UpdateUserContact(ctx, models.UserContact{
		Id:          userContactId,
		UserId:      userId,
		Email:       userContact.Email,
		Name:        userContact.Name,
		PrefixPhone: userContact.PrefixPhone,
		Phone:       userContact.Phone,
		UpdatedAt:   date.ContactDate(time.Now()),
	})

	if err != nil {
		return userContactDB, errors.NewUserErrorWithError("user_contact.error", http.StatusBadRequest, err)
	}

	go func() {
		_, err := u.userContactModificationRepository.CreateModification(context.Background(), userContactDB)
		if err != nil {
			fmt.Printf("Error while creating modification on update %v for user: %d", err, userId)
		}
	}()

	return userContactDB, nil
}

func (u *userContactService) Delete(ctx context.Context, userId, userContactId int64) (models.UserContact, error) {

	userContact, err := u.userContactRepository.FindUserContactByUserIdAndId(ctx, userId, userContactId)

	if err != nil {
		return models.UserContact{}, errors.NewUserErrorWithError("user_contact.not_found", http.StatusNotFound, err)
	}

	contact, err := u.userContactRepository.DeleteUserContact(ctx, userContact)

	if err != nil {
		return contact, errors.NewUserErrorWithError("user_contact.error", http.StatusBadRequest, err)
	}

	go func() {
		_, err := u.userContactModificationRepository.CreateModification(context.Background(), contact)
		if err != nil {
			fmt.Printf("Error while creating modification on delete %v for user: %d", err, userId)
		}
	}()

	return contact, nil
}

func (u *userContactService) FindAllByUserId(ctx context.Context, userId int64) ([]dto.BasicUserContact, error) {

	return u.userContactRepository.FindAllUserContactByUserId(ctx, userId)
}

func (u *userContactService) FindByUserIdAndId(context context.Context, userId, userContactId int64) (models.UserContact, error) {

	userContact, err := u.userContactRepository.FindUserContactByUserIdAndId(context, userId, userContactId)

	if err != nil {
		return userContact, errors.NewUserErrorWithError("user_contact.error", http.StatusBadRequest, err)
	}

	if userContact.Id == 0 {
		return userContact, errors.NewUserError("user_contact.not_found", http.StatusNotFound)
	}

	return userContact, nil
}

func ProvideUserContactService(
	userContactRepository repositories.UserContactRepository,
	userContactModificationRepository repositories.UserContactModificationRepository,
) UserContactService {

	return &userContactService{
		userContactRepository:             userContactRepository,
		userContactModificationRepository: userContactModificationRepository,
	}

}
