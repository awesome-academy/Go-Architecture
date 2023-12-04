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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.POST("/task", tc.Create)
	group.GET("/task", tc.Fetch)
}
