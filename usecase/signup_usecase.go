package usecase

import (
	"Go-Architecture/domain"
	"Go-Architecture/domain/entities"
	"Go-Architecture/repository"
	"golang.org/x/net/context"
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

func (s signupUsecase) GetUserByEmail(context context.Context, email string) (entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s signupUsecase) Create(context context.Context, user *entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (s signupUsecase) CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error) {
	//TODO implement me
	panic("implement me")
}

func (s signupUsecase) CreateRefreshToken(user *entities.User, secret string, expiry int) (refreshToken string, err error) {
	//TODO implement me
	panic("implement me")
}
