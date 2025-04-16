package routes

import (
	"database/sql"
	"net/http"

	"github.com/eralves01/user-services/internal/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(handler *handlers.UserHandler, db *sql.DB, router *gin.RouterGroup) {
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Lista de usuários"})
	})

	router.POST("/user", handler.Create)

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"message": "Detalhes do usuário", "id": id})
	})
}
