package controller

import (
	"Go-Architecture/bootstrap"
	"Go-Architecture/domain"
	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct {
	RefreshTokenUsecase domain.RefreshTokenUsecase
	Env                 *bootstrap.Env
}

func (rtc RefreshTokenController) RefreshToken(context *gin.Context) {
	// Declare request

	// Extra userId from token

	// get User by ID

	// Create acess token

	// Create refresh token
}
