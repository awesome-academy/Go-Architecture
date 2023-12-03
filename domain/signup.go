package domain

import (
	"Go-Architecture/domain/entities"
	"golang.org/x/net/context"
)

type SignupRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"require"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	GetUserByEmail(context context.Context, email string) (entities.User, error)
	Create(context context.Context, user *entities.User) error
	CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *entities.User, secret string, expiry int) (refreshToken string, err error)
}
