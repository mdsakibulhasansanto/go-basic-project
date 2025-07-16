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

func (r *MySQLProductRepository) Update(product *models.Product) error {
	result := r.db.Model(&models.Product{}).Where("id = ?", product.ID).Updates(product)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

// Delete deletes a product by its ID
func (r *MySQLProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
