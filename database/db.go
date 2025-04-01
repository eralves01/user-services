package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var instance *sql.DB
var once sync.Once

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) GetInstance() *sql.DB {
	once.Do(func() {
		log.Println("Connecting to database...")
		var err error
		instance, err = connect()
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
	})
	return instance
}

func connect() (*sql.DB, error) {
	dbURL, err := createDatabaseURL()
	if err != nil {
		return nil, err
	}
	return sql.Open("postgres", dbURL)
}

func createDatabaseURL() (string, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		return "", fmt.Errorf("missing one or more required database environment variables")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName), nil
}

// CloseConnection fecha a conexão com o banco
func (db *Database) CloseConnection() {
	if instance != nil {
		err := instance.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed.")
		}
	}
}
