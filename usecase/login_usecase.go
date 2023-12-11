package usecase

import (
	"Go-Architecture/domain"
	"Go-Architecture/domain/entity"
	"Go-Architecture/repository"
	"Go-Architecture/utils/tokenutil"
	"context"
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

func (lu loginUsecase) GetUserByEmail(c context.Context, email string) (entity.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu loginUsecase) CreateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu loginUsecase) CreateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
