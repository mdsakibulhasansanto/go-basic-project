package repositories

import "final-project/models"

type UserRepositoryInterface interface {
	Create(user models.User) error
	GetByEmail(email string) (*models.User, error)
	UpdateUserByID(user *models.User) error
	DeleteUser(email string) error
}
