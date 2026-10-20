package domain

import (
	"context"
)

type Task struct {
	ID     int64  `gorm:"primaryKey;column:id" json:"id"`
	Title  string `gorm:"column:title" json:"title" form:"title" binding:"required"`
	UserID int64  `gorm:"column:user_id" json:"user_id"`
}

func (Task) TableName() string {
	return "tasks"
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID int64) ([]Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID int64) ([]Task, error)
}
