package main

import (
	"fmt"
	"net/http"

	"github.com/eralves01/user-services/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Rodando")
	router := gin.Default()

	routes.RegisterRoutes(router)

	http.ListenAndServe(":8080", router)
}
