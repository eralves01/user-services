package repositories

import (
	"database/sql"

	"github.com/eralves01/user-services/internal/domain"
)

type UserTypeRepository struct {
	db *sql.DB
}

func NewUserTypeRepository(db *sql.DB) *UserTypeRepository {
	return &UserTypeRepository{db: db}
}

func (u UserTypeRepository) Create(id int, value string) error {
	stmt, err := u.db.Prepare("INSERT INTO user_types (id, value) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id, value); err != nil {
		return err
	}

	return nil
}

func (u UserTypeRepository) GetUserTypeByValue(value string) (*domain.UserType, error) {
	var userType domain.UserType
	err := u.db.QueryRow("SELECT id, value FROM user_types WHERE value=$1", value).Scan(&userType.ID, &userType.Value)
	if err == sql.ErrNoRows {
		return nil, domain.ErrUserTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	return &userType, nil
}

func (u UserTypeRepository) GetUserTypeByID(id int) (*domain.UserType, error) {
	var userType domain.UserType
	err := u.db.QueryRow("SELECT id, value FROM user_types WHERE id=$1", id).Scan(&userType.ID, &userType.Value)
	if err == sql.ErrNoRows {
		return nil, domain.ErrUserTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	return &userType, nil
}

func (u UserTypeRepository) GetUserTypes() ([]domain.UserType, error) {
	rows, err := u.db.Query("SELECT id, value FROM user_types")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userTypes []domain.UserType
	for rows.Next() {
		var userType domain.UserType
		if err := rows.Scan(&userType.ID, &userType.Value); err != nil {
			return nil, err
		}
		userTypes = append(userTypes, userType)
	}

	return userTypes, nil
}
