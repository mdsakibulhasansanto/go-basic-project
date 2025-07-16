package services

import (
	"errors"
	"final-project/models"
	"final-project/repositories"
)

type ProductService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductService) Update(product *models.Product) error {

	return s.repo.Update(product)
}

// Delete removes a product by its ID.
// Delete removes a product by its ID.
func (s *ProductService) Delete(id uint) error {
	if id == 0 {
		return errors.New("invalid product ID")
	}
	return s.repo.Delete(id)
}
