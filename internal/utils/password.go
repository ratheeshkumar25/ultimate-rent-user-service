package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword will hasj the password and return the hash password
func HashPassword(pasword string) (string, error) {
	// GenerateFromPassword returns the bcrypt hash of the password at the given cost
	bytes, err := bcrypt.GenerateFromPassword([]byte(pasword), 14)
	if err != nil {
		return "", err
	}
	Password := string(bytes)
	return Password, nil
}

// CheckPassword function will check the provide password with userPassword
func CheckPassword(providedPassword, password string) bool {
	// CompareHashAndPassword compares a bcrypt hashed password with its possible plaintext equivalent
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	return err == nil
}
