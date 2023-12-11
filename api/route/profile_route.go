package route

import (
	"Go-Architecture/api/controller"
	"Go-Architecture/repository"
	"Go-Architecture/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func NewProfileRouter(timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
	group.PUT("/profile", pc.Update)
}
