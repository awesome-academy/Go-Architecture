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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)

	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}

	group.POST("/signup", sc.Signup)
}
