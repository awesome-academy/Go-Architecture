package domain

import (
	"Go-Architecture/domain/entity"
	"context"
)

type SignupRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
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
