package database

import (
	"clean_arc/features/users/repository"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(repository.User{})
}