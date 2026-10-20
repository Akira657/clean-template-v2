package route

import (
	"time"

	"clean-template/bootstrap"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewHealthRouter(env, timeout, publicRouter)

	//protectedRouter := gin.Group("")
	//// Middleware to verify AccessToken
	//protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	//// All Private APIs
	//NewProfileRouter(env, timeout, db, protectedRouter)
	//NewTaskRouter(env, timeout, db, protectedRouter)
}
