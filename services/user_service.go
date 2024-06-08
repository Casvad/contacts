package services

import (
	"contacts/dto"
	"contacts/models"
	"contacts/repositories"
	"contacts/utils"
	"contacts/utils/date"
	"contacts/utils/env"
	"contacts/utils/errors"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type userService struct {
	userRepository repositories.UserRepository
}

func (u *userService) Register(context context.Context, user dto.RegisterUser) (models.User, error) {

	userDB, err := u.userRepository.FindUserByEmail(context, user.Email)
	if err != nil {
		return models.User{}, errors.NewUserError("user.error", http.StatusInternalServerError)
	}
	if userDB.Id != 0 {
		return models.User{}, errors.NewUserError("user.already.exists", http.StatusConflict)
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, errors.NewUserErrorWithError("user.bad.password", http.StatusBadRequest, err)
	}

	return u.userRepository.CreateUser(context, models.User{
		Email:     user.Email,
		Password:  hashPassword,
		Name:      user.Name,
		CreatedAt: date.ContactDate(time.Now()),
		UpdatedAt: date.ContactDate(time.Now()),
	})
}

func (u *userService) Login(context context.Context, login dto.Login) (dto.LoginResponse, error) {
	user, err := u.userRepository.FindUserByEmail(context, login.Email)
	if err != nil {
		return dto.LoginResponse{}, errors.NewUserError("user.not_found", http.StatusNotFound)
	}

	if !utils.CheckPasswordHash(login.Password, user.Password) {
		return dto.LoginResponse{}, errors.NewUserError("user.wrong_password", http.StatusBadRequest)
	}

	expirationTime := time.Now().Add(50 * time.Minute)
	claims := &dto.Claims{
		Email:  user.Email,
		UserId: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenString, err := token.SignedString([]byte(env.JwtKey))
	if err != nil {
		return dto.LoginResponse{}, errors.NewUserErrorWithError("user.bad_login", http.StatusInternalServerError, err)
	}

	return dto.LoginResponse{
		Token: tokenString,
	}, nil
}

func ProvideUserService(userRepository repositories.UserRepository) UserService {

	return &userService{userRepository}
}
