package repositories

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/eralves01/user-services/dto"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	_ = godotenv.Load(".env.test") // Carrega config de testes

	var err error
	testDB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	code := m.Run()
	testDB.Close()
	os.Exit(code)
}

func TestCreateUser_Integration(t *testing.T) {
	repo := NewUserRepository(testDB)

	input := dto.CreateUserDTO{
		Name:       "Integration User",
		Email:      "integration@example.com",
		UserTypeID: 1,
		Password:   "hash123",
	}

	result, err := repo.Create(input)
	assert.NoError(t, err)
	assert.Equal(t, input.Email, result.Email)
	assert.Equal(t, input.UserTypeID, result.UserTypeID)
}
