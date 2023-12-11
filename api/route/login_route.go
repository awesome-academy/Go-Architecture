package route

import (
	"Go-Architecture/api/controller"
	"Go-Architecture/bootstrap"
	"Go-Architecture/repository"
	"Go-Architecture/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
