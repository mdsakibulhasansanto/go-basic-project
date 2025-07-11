package repositories

import (
	"final-project/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
}

type MySQLProductRepository struct {
	db *gorm.DB
}

func NewMySQLProductRepository(db *gorm.DB) ProductRepository {
	return &MySQLProductRepository{db: db}
}

func (r *MySQLProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}
