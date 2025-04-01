package repositories

import (
	"database/sql"
	"log"

	"github.com/eralves01/user-services/dto"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user dto.CreateUserDTO) error {
	log.Print(user)
	return nil
}
