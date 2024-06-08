package services

import (
	"contacts/dto"
	"contacts/jobs"
	"contacts/models"
	"contacts/repositories"
	"contacts/utils/env"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserContactService_Create(t *testing.T) {
	gorm := repositories.ProvideGormFromEnv()
	CleanDatabase(gorm)
	env.MigrationPath = "file://../db/migrations"
	job := jobs.ProvideMigrationJob()
	job.Execute()
	userRepository := repositories.ProvideUserRepository(gorm)
	userContactService := ProvideUserContactService(repositories.ProvideUserContactRepository(gorm), repositories.ProvideUserContactModificationRepository(gorm))

	user, _ := userRepository.CreateUser(context.Background(), models.User{
		Email:    "x@gmail.com",
		Password: "x123",
		Name:     "carlos",
	})

	uc, err := userContactService.Create(context.Background(), user.Id, dto.UserContact{
		Email:       "x2@gmail.com",
		Name:        "X",
		PrefixPhone: "+57",
		Phone:       "345353421",
	})

	assert.Nil(t, err)
	assert.True(t, uc.Id > 0)
	assert.Equal(t, uc.UserId, user.Id)
}
