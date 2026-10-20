package domain

import (
	"context"
)

type User struct {
	ID           int64  `gorm:"primaryKey;column:id" json:"id"`
	Username     string `gorm:"column:username;unique;not null" json:"username"`
	PasswordHash string `gorm:"column:password_hash;not null" json:"-"`
	Salt         string `gorm:"column:salt;not null" json:"-"`
	IsSuperuser  bool   `gorm:"column:is_superuser;default:false" json:"is_superuser"`
	CreatedAt    int64  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt    int64  `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByUsername(c context.Context, username string) (User, error)
	GetByID(c context.Context, id int64) (User, error)
}
