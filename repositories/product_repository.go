package repositories

import (
	"final-project/models"

	"gorm.io/gorm"
)

type MySQLProductRepository struct {
	db *gorm.DB
}

func NewMySQLProductRepository(db *gorm.DB) ProductRepository {
	return &MySQLProductRepository{db: db}
}

func (r *MySQLProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *MySQLProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
