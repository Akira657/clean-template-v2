package repository

import (
	"context"

	"clean-template/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	return ur.db.WithContext(c).Create(user).Error
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	var users []domain.User
	err := ur.db.WithContext(c).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetByUsername(c context.Context, username string) (domain.User, error) {
	var user domain.User
	err := ur.db.WithContext(c).Where("username = ?", username).First(&user).Error
	return user, err
}

func (ur *userRepository) GetByID(c context.Context, id int64) (domain.User, error) {
	var user domain.User
	err := ur.db.WithContext(c).Where("id = ?", id).First(&user).Error
	return user, err
}
