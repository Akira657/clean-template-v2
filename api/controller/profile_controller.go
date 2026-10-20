package controller

import (
	"net/http"
	"strconv"

	"clean-template/domain"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

func (pc *ProfileController) Fetch(c *gin.Context) {
	userIDStr := c.GetString("x-user-id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid user ID"})
		return
	}

	profile, err := pc.ProfileUsecase.GetProfileByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
