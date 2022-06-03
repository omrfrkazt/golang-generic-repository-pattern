package db

import (
	"github.com/omrfrkazt/golang-generic-repository-pattern/internal/entities"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&entities.User{})
}
