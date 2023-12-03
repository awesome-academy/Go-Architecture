package usecase

import (
	"Go-Architecture/domain"
	"Go-Architecture/repository"
	"time"
)

type loginUsecase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository repository.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
