package route

import (
	"Go-Architecture/bootstrap"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewSignupRouter(env, timeout, db, publicRouter)
}
