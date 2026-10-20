package postgres

import "gorm.io/gorm"

type PostgresDatabase struct {
	DB *gorm.DB
}
