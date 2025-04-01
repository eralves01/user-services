package services

import (
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

func (s *UserService) Create(user dto.CreateUserDTO) error {
	passManager := utils.NewPasswordManager()

	passwordHash, err := passManager.GetPasswordHash(user.Password)
	if err != nil {
		log.Printf("Error generating password hash: %v", err)
		return err
	}

	user.PasswordHash = string(passwordHash)

	if err := s.repository.Create(user); err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}

	return nil
}
