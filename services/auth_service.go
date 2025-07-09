package services

import (
	"errors"
	"final-project/models"
	"final-project/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repositories.UserRepositoryInterface
}

func NewAuthService(repo repositories.UserRepositoryInterface) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(username, email, password string) error {
	existingUser, _ := s.repo.GetByEmail(email)
	if existingUser != nil {
		return errors.New("user already exists with this email")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username:          username,
		Email:             email,
		PasswordHash:      string(hashedPassword),
		IsVerified:        false,
		VerificationToken: "token123",
	}

	return s.repo.Create(user)
}

func (s *AuthService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetByEmail(email)
}
