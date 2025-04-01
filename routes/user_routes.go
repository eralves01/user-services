package routes

import (
	"database/sql"
	"net/http"

	"github.com/eralves01/user-services/handlers"
	"github.com/eralves01/user-services/repositories"
	"github.com/eralves01/user-services/services"
	"github.com/gin-gonic/gin"
)

func UserRoutes(db *sql.DB, router *gin.RouterGroup) {
	repository := repositories.NewUserRepository(db)
	service := services.NewUserService(repository)
	handler := handlers.NewUserHandler(service)

	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Lista de usuários"})
	})

	router.POST("/user", handler.Create)

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"message": "Detalhes do usuário", "id": id})
	})
}
