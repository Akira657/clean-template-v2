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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
