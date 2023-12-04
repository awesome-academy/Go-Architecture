package usecase

import (
	"Go-Architecture/domain"
	"Go-Architecture/domain/entity"
	"Go-Architecture/repository"
	"context"
	"time"
)

type profileUsecase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository repository.UserRepository, timeout time.Duration) domain.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string) (*entity.UserProfile, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &entity.UserProfile{Name: user.Name, Email: user.Email}, nil
}
