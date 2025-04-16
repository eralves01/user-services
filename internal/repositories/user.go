package repositories

import (
	"database/sql"

	"github.com/eralves01/user-services/internal/domain"
	interfaces "github.com/eralves01/user-services/internal/domain/interfaces"
	"github.com/google/uuid"
)

type UserRepository struct {
	db interfaces.DBExecutor
}

func NewUserRepository(db interfaces.DBExecutor) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user domain.User) error {
	stmt, err := r.db.Prepare("INSERT INTO users (name, email, user_type_id, password_hash) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(
		user.Name,
		user.Email,
		user.UserTypeID,
		user.PasswordHash,
	); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, name, email, user_type_id, password_hash FROM users WHERE email=$1"
	err := r.db.QueryRow(
		query, email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.UserTypeID, &user.PasswordHash)
	if err == sql.ErrNoRows {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, name, email, user_type_id, password_hash FROM users WHERE id=$1"
	err := r.db.QueryRow(
		query, id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.UserTypeID, &user.PasswordHash)
	if err == sql.ErrNoRows {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	query := "SELECT id, name, email, user_type_id, password_hash FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.UserTypeID, &user.PasswordHash); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
