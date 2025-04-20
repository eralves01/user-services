package domain

import "github.com/eralves01/user-services/internal/domain"

type UserTypeRepository interface {
	Create(userType domain.UserType) error
	GetByID(id int) (*domain.UserType, error)
	GetByValue(value string) (*domain.UserType, error)
	GetAll() ([]domain.UserType, error)
}
