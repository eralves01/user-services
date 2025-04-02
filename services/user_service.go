package services

import (
	"errors"
	"log"

	"github.com/eralves01/user-services/dto"
	"github.com/eralves01/user-services/repositories"
	"github.com/eralves01/user-services/utils"
)

type UserService struct {
	repository repositories.UserRepositoryInterface
}

func NewUserService(repository repositories.UserRepositoryInterface) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) Create(user dto.CreateUserDTO) (*dto.UserResponseDTO, error) {
	passManager := utils.NewPasswordManager()

	passwordHash, err := passManager.GetPasswordHash(user.Password)
	if err != nil {
		log.Printf("Error generating password hash: %v", err)
		return nil, err
	}

	user.PasswordHash = string(passwordHash)

	userTypeID, err := s.repository.GetUserTypeID(user.UserType)
	if err != nil || userTypeID == -1 {
		log.Print("User type is not found!")
		return nil, errors.New("User type is not found!")
	}

	user.UserTypeID = userTypeID

	userResponse, err := s.repository.Create(user)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return nil, err
	}

	return userResponse, nil
}
