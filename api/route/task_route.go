package route

import (
	"time"

	"clean-template/api/controller"
	"clean-template/bootstrap"
	"clean-template/repository"
	"clean-template/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
