package services

import (
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
func (s *ProductService) Delete(id uint) error {
	// TODO: Implement actual deletion logic, e.g., using a repository or database.
	// For now, return nil or an error if not found.
	// Example:
	// err := s.repo.DeleteByID(id)
	// if err != nil {
	//     return err
	// }
	// return nil

	// Placeholder implementation:
	return nil
}
