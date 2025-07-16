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

// UpdateUserByID updates user data by ID
func (s *AuthService) UpdateUserByID(user *models.User) error {
	if user.ID == 0 {
		return errors.New("user ID is required for update")
	}
	return s.repo.UpdateUserByID(user)
}

// DeleteUser deletes a user by email
func (s *AuthService) DeleteUser(email string) error {
	if email == "" {
		return errors.New("email is required for deletion")
	}
	return s.repo.DeleteUser(email)
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
		VerificationToken: "123token",
	}

	err = s.repo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func (s *AuthService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetByEmail(email)
}
