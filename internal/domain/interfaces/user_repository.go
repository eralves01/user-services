package domain

import "github.com/eralves01/user-services/internal/domain"

type UserRepository interface {
	Create(user domain.User) error
}
