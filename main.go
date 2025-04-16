package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eralves01/user-services/configs"
	"github.com/eralves01/user-services/database"
	"github.com/eralves01/user-services/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfig()

	db := database.NewDatabase().GetInstance()
	defer database.NewDatabase().CloseConnection()

	err := db.Ping()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	} else {
		log.Println("Database connected successfully!")
	}

	if os.Getenv("RUN_MIGRATIONS") == "true" {
		database.RunMigrations()
	}

	router := gin.Default()

	r := routes.NewRegisterRoutes(db, router)
	r.GetRegisterRoutes()

	http.ListenAndServe(":8080", router)
}
