package route

import (
	"Go-Architecture/api/middleware"
	"Go-Architecture/bootstrap"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All privates APIs
	// NewProfileRouter
}
