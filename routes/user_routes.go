package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Lista de usuários"})
	})

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"message": "Detalhes do usuário", "id": id})
	})
}
