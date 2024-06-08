package dto

import (
	"github.com/golang-jwt/jwt/v4"
)

type BasicUserContact struct {
	Id    int64  `json:"id" gorm:"primaryKey"`
	Email string `gorm:"column:email" json:"email"`
	Name  string `gorm:"column:name" json:"name"`
}

type RegisterUser struct {
	Email    string `gorm:"column:email" json:"email"`
	Name     string `gorm:"column:name" json:"name"`
	Password string `gorm:"column:password" json:"password"`
}
type Login struct {
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserContact struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	PrefixPhone string `json:"prefix_phone"`
	Phone       string `json:"phone"`
}

type Claims struct {
	Email  string `json:"email"`
	UserId int64  `json:"user_id"`
	jwt.RegisteredClaims
}
