package domain

import (
	"Go-Architecture/domain/entity"
	"context"
)

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID string) (*entity.UserProfile, error)
}
