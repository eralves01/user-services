package database

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	log.Print("Applying migrations")
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
