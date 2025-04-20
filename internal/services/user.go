package services

import (
	interfaces "github.com/eralves01/user-services/internal/domain/interfaces"
	"github.com/eralves01/user-services/internal/dto"
)

type UserService struct {
	repository      interfaces.UserRepository
	userTypeService *UserTypeService
}

func NewUserService(repository interfaces.UserRepository, userTypeService *UserTypeService) *UserService {
	return &UserService{
		repository:      repository,
		userTypeService: userTypeService,
	}
}

func (s *UserService) Create(user dto.CreateUserInput) error {
	_, err := s.userTypeService.reposytory.GetByID(user.UserTypeID)
	if err != nil {
		return err
	}

	if err := user.Validate(); err != nil {
		return err
	}

	newUser := user.ToUser()

	return s.repository.Create(*newUser)
}
