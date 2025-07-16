package repositories

import "final-project/models"

type ProductRepository interface {
	Create(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
}
