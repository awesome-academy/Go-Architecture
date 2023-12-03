package usecase

import (
	"Go-Architecture/domain"
	"Go-Architecture/repository"
	"time"
)

type refreshTokenUsecase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUsecase(userRepository repository.UserRepository, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
