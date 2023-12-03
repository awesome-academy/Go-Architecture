package controller

import (
	"Go-Architecture/bootstrap"
	"Go-Architecture/domain"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

func (lc LoginController) Login(context *gin.Context) {

}
