package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordManager struct{}

func NewPasswordManager() *PasswordManager {
	return &PasswordManager{}
}

func (p *PasswordManager) GetPasswordHash(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return hashedPassword, err
	}

	return hashedPassword, nil
}

func (p *PasswordManager) CompareHashAndPassword(hashedPassword []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	return err == nil
}
