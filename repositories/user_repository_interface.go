package repositories

import "github.com/eralves01/user-services/dto"

type UserRepositoryInterface interface {
	Create(user dto.CreateUserDTO) error
}
