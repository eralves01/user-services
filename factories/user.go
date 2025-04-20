package factories

import (
	"database/sql"
	"log"

	"github.com/eralves01/user-services/internal/dto"
	"github.com/google/uuid"
)

type NewUser struct {
	ID           uuid.UUID
	Name         string
	Email        string
	UserType     int
	PasswordHash string
}

func CreateNewUser(db *sql.DB, user dto.CreateUserInput) NewUser {
	prepareUser(db, &user)
	insertedUser := createUser(db, user)

	return insertedUser
}

func prepareUser(db *sql.DB, user *dto.CreateUserInput) *dto.CreateUserInput {
	if user.Name == "" {
		user.Name = "Test User"
	}
	if user.Email == "" {
		user.Email = generateEmail()
	}
	if user.Password == "" {
		user.Password = "123456"
	}

	return user
}

func generateEmail() string {
	return "user_" + uuid.NewString()[:8] + "@test.com"
}

func createUser(db *sql.DB, user dto.CreateUserInput) NewUser {
	var newUser NewUser
	err := db.QueryRow(`
		INSERT INTO users (name, email, type, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, email, type, password_hash
	`, user.Name, user.Email, user.UserTypeID, user.Password).
		Scan(
			&newUser.ID,
			&newUser.Name,
			&newUser.Email,
			&newUser.UserType,
			&newUser.PasswordHash,
		)

	if err != nil {
		log.Fatalf("Erro ao criar user %v", err)
	}

	return newUser
}
