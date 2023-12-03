package domain

import (
	"Go-Architecture/domain/entity"
	"context"
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
	GetUserByEmail(context context.Context, email string) (entity.User, error)
	Create(context context.Context, user *entity.User) error
	CreateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error)
}
