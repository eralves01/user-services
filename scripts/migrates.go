package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eralves01/user-services/configs"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	runMigrations()
}

func runMigrations() {
	log.Print("Applying migrations")

	configs.LoadConfig()
	databaseURL, err := createDatabaseURL()
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.New(
		"file://database/migrations",
		databaseURL,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Print("Migrations successfully implemented!")
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
