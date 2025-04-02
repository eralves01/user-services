package repositories

import (
	"github.com/eralves01/user-services/constants"
	"github.com/eralves01/user-services/dto"
)

type UserRepositoryInterface interface {
	Create(user dto.CreateUserDTO) (*dto.UserResponseDTO, error)
	GetUserTypeID(userType constants.UserType) (int, error)
}
