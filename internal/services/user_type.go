package services

import (
	"github.com/eralves01/user-services/internal/domain"
	interfaces "github.com/eralves01/user-services/internal/domain/interfaces"
)

type UserTypeService struct {
	reposytory interfaces.UserTypeRepository
}

func NewUserTypeService(reposytory interfaces.UserTypeRepository) *UserTypeService {
	return &UserTypeService{reposytory: reposytory}
}

func (s *UserTypeService) GetUserTypeByValue(value string) (*domain.UserType, error) {
	userType, err := s.reposytory.GetUserTypeByValue(value)
	if err != nil {
		return nil, err
	}

	return userType, nil
}
