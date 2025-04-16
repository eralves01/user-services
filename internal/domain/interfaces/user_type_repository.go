package domain

import "github.com/eralves01/user-services/internal/domain"

type UserTypeRepository interface {
	Create(id int, value string) error
	GetUserTypeByID(id int) (*domain.UserType, error)
	GetUserTypeByValue(value string) (*domain.UserType, error)
}
