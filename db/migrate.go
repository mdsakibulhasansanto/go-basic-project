package db

import (
	"final-project/models"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {

	return db.AutoMigrate(
		&models.Product{},
		&models.User{},
	)
}
