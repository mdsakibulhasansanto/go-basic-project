package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model  // Includes ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string
	Description string
	Price       float64
	UserEmail   string
}
