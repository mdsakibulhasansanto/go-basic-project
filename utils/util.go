package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func SendVerificationEmail(email, verificatinToken string) {

	fmt.Printf("Send verification token %s to email %s", verificatinToken, email)
}
