package usecase

import (
	"Go-Architecture/domain"
	"Go-Architecture/domain/entity"
	"Go-Architecture/repository"
	"context"
	"time"
)

type signupUsecase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository repository.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su signupUsecase) GetUserByEmail(c context.Context, email string) (entity.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func (s signupUsecase) Create(c context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (s signupUsecase) CreateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error) {
	//TODO implement me
	panic("implement me")
}

func (s signupUsecase) CreateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error) {
	//TODO implement me
	panic("implement me")
}
