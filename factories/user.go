package factories

import (
	"database/sql"
	"log"

	"github.com/eralves01/user-services/constants"
	"github.com/eralves01/user-services/dto"
	"github.com/eralves01/user-services/utils"
	"github.com/google/uuid"
)

type NewUser struct {
	ID           uuid.UUID
	Name         string
	Email        string
	UserType     int
	PasswordHash string
}

func CreateNewUser(db *sql.DB, user dto.CreateUserDTO) NewUser {
	prepareUser(db, &user)
	insertedUser := createUser(db, user)

	return insertedUser
}

func prepareUser(db *sql.DB, user *dto.CreateUserDTO) *dto.CreateUserDTO {
	if user.Name == "" {
		user.Name = "Test User"
	}
	if user.Email == "" {
		user.Email = generateEmail()
	}
	if user.PasswordHash == "" {
		pass := utils.NewPasswordManager()
		passwordHash, err := pass.GetPasswordHash("password")
		if err != nil {
			log.Fatal(err)
		}
		user.PasswordHash = string(passwordHash)
	}
	if user.UserType == constants.Client {
		CreateUserTypeClient(db)
	} else {
		CreateUserTypeMerchant(db)
	}

	return user
}

func generateEmail() string {
	return "user_" + uuid.NewString()[:8] + "@test.com"
}

func createUser(db *sql.DB, user dto.CreateUserDTO) NewUser {
	var newUser NewUser
	err := db.QueryRow(`
		INSERT INTO users (name, email, type, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, email, type, password_hash
	`, user.Name, user.Email, user.UserType, user.PasswordHash).
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
