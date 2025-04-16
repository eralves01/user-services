package domain

import (
	"log"
	"time"

	"github.com/eralves01/user-services/internal/utils"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Email        string
	UserTypeID   int
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser(name string, email string, userType int, password string) *User {
	return &User{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		UserTypeID:   userType,
		PasswordHash: encryptPassword(password),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func encryptPassword(password string) string {
	passwordHash, err := utils.GetPasswordHash(password)
	if err != nil {
		log.Fatal(err)
	}

	return string(passwordHash)
}
