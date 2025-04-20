package dto

import "time"

type UserOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserOutput(id, name, email, userType string, createdAt, updatedAt time.Time) *UserOutput {
	return &UserOutput{
		ID:        id,
		Name:      name,
		Email:     email,
		UserType:  userType,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
