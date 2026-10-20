package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func (hc *HealthController) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
