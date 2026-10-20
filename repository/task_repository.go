package repository

import (
	"context"

	"clean-template/domain"

	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	return tr.db.WithContext(c).Create(task).Error
}

func (tr *taskRepository) FetchByUserID(c context.Context, userID int64) ([]domain.Task, error) {
	var tasks []domain.Task
	err := tr.db.WithContext(c).Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
