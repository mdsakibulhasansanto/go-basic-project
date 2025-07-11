package repositories

import "final-project/models"

type ProductRepository interface {
	Create(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
}
