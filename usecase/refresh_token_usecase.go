package usecase

import (
	"Go-Architecture/domain"
	"Go-Architecture/domain/entity"
	"Go-Architecture/repository"
	"Go-Architecture/utils/tokenutil"
	"context"
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

func (rtu refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}

func (rtu refreshTokenUsecase) GetUserByID(c context.Context, id string) (entity.User, error) {
	ctx, cancel := context.WithTimeout(c, rtu.contextTimeout)
	defer cancel()
	return rtu.userRepository.GetByID(ctx, id)
}

func (rtu refreshTokenUsecase) CreateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (rtu refreshTokenUsecase) CreateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
