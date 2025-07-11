package db

import (
	"final-project/models"

	"gorm.io/gorm"
)


func MigrateProductTable (db *gorm.DB) error {

	return db.AutoMigrate(&models.Product{})
}