package factories

import (
	"database/sql"
	"log"

	"github.com/eralves01/user-services/constants"
)

func CreateUserTypeClient(db *sql.DB) int {
	return createUserType(db, string(constants.Client))
}

func CreateUserTypeMerchant(db *sql.DB) int {
	return createUserType(db, string(constants.Merchant))
}

func createUserType(db *sql.DB, userType string) int {
	var id int
	err := db.QueryRow(`
		INSERT INTO user_types (type)
		VALUES ($1)
		RETURNING id
	`, userType).Scan(&id)

	if err != nil {
		log.Fatalf("Erro ao criar user_type '%s': %v", userType, err)
	}

	return id
}
