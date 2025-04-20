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

func (s *UserTypeService) Create(userType domain.UserType) error {
	if userTypeReturn, _ := s.reposytory.GetByValue(userType.Value); userTypeReturn != nil {
		return domain.ErrUserTypeAlreadyExists
	}
	if userTypeReturn, _ := s.reposytory.GetByID(userType.ID); userTypeReturn != nil {
		return domain.ErrUserTypeAlreadyExists
	}

	return s.reposytory.Create(userType)
}

func (s *UserTypeService) GetByValue(value string) (*domain.UserType, error) {
	userType, err := s.reposytory.GetByValue(value)
	if err != nil {
		return nil, err
	}

	return userType, nil
}

func (s *UserTypeService) GetByID(id int) (*domain.UserType, error) {
	userType, err := s.reposytory.GetByID(id)
	if err != nil {
		return nil, err
	}

	return userType, nil
}

func (s *UserTypeService) GetAll() ([]domain.UserType, error) {
	userTypes, err := s.reposytory.GetAll()
	if err != nil {
		return nil, err
	}

	return userTypes, nil
}
