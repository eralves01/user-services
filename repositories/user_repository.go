package repositories

import (
	"log"

	"github.com/eralves01/user-services/constants"
	"github.com/eralves01/user-services/dto"
)

type UserRepository struct {
	db DBExecutor
}

func NewUserRepository(db DBExecutor) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user dto.CreateUserDTO) (*dto.UserResponseDTO, error) {
	var userResponse dto.UserResponseDTO
	err := r.db.QueryRow(
		"INSERT INTO users (name, email, user_type_id, password_hash) VALUES ($1, $2, $3, $4) RETURNING id, name, email, user_type_id",
		user.Name, user.Email, user.UserTypeID, user.PasswordHash).
		Scan(&userResponse.ID, &userResponse.Name, &userResponse.Email, &userResponse.UserTypeID)

	if err != nil {
		log.Fatal("Erro ao inserir usuário:", err)
		return nil, err
	}
	return &userResponse, nil
}

func (r *UserRepository) GetUserTypeID(userType constants.UserType) (int, error) {
	var id int
	query := "SELECT id FROM user_types WHERE type=$1"
	err := r.db.QueryRow(
		query, userType,
	).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
