package domain

import (
	"Go-Architecture/domain/entity"
	"context"
)

type UpdateProfileRequest struct {
	UserName string `form:"userName"`
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID string) (*entity.UserProfile, error)
	UpdateUserName(c context.Context, userID string, userName string) error
}
