package route

import (
	"clean-template/api/controller"
	"clean-template/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func NewHealthRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	hc := &controller.HealthController{}
	group.GET("/health", hc.Health)
}
