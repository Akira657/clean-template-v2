package controller

import (
	"net/http"
	"strconv"

	"clean-template/domain"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc *TaskController) Create(c *gin.Context) {
	var task domain.Task

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userIDStr := c.GetString("x-user-id")
	task.UserID, err = strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid user ID"})
		return
	}

	err = tc.TaskUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (u *TaskController) Fetch(c *gin.Context) {
	userIDStr := c.GetString("x-user-id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid user ID"})
		return
	}

	tasks, err := u.TaskUsecase.FetchByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
