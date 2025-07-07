package services

import (
	"final-project/models"
	"final-project/repositories"
	"final-project/utils"
	"fmt"

	"github.com/google/uuid"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) RegisterUser(username, email, password string) error {
	// ✅ Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	// ✅ Generate verification token
	verificationToken := uuid.New().String()

	// ✅ Check if user already exists
	existingUser, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return fmt.Errorf("error checking existing user: %v", err)
	}
	if existingUser != nil {
		return fmt.Errorf("user already exists with email: %s", email)
	}

	// ✅ Create user model
	user := models.User{
		Username:          username,
		Email:             email,
		PasswordHash:      hashedPassword,
		IsVerified:        false,
		VerificationToken: verificationToken,
	}

	// ✅ Save to database
	if err := s.userRepo.Create(user); err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (s *AuthService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetByEmail(email)
}
