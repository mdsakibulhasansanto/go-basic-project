package repositories

import (
	"database/sql"
	"final-project/models"
	"fmt"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Create(user models.User) error {
	query := "INSERT INTO users (username, email, password, is_verified, verification_token) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, user.Username, user.Email, user.PasswordHash, user.IsVerified, user.VerificationToken)
	if err != nil {
		return fmt.Errorf("error during user creation: %v", err)
	}
	return nil
}

func (r *MySQLUserRepository) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id, username, email, password, is_verified, verification_token FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.IsVerified, &user.VerificationToken)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
