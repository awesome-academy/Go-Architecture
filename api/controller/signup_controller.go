package controller

import (
	"Go-Architecture/bootstrap"
	"Go-Architecture/domain"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}
