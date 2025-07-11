package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySqlGormDB() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/go_test_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error opening database:", err)
		return nil, err
	}

	fmt.Println("Database connection successful.")
	return db, nil
}
