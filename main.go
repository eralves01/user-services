package main

import (
	"log"
	"net/http"

	"github.com/eralves01/user-services/configs"
	"github.com/eralves01/user-services/database"
	"github.com/eralves01/user-services/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfig()

	// Criar instância do banco
	db := database.NewDatabase().GetInstance()
	defer database.NewDatabase().CloseConnection()

	// Testar conexão
	err := db.Ping()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	} else {
		log.Println("Database connected successfully!")
	}

	router := gin.Default()

	r := routes.NewRegisterRoutes(db, router)
	r.GetRegisterRoutes()

	http.ListenAndServe(":8080", router)
}
