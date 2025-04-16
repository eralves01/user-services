package routes

import (
	"database/sql"
	"net/http"

	"github.com/eralves01/user-services/internal/handlers"
	"github.com/gin-gonic/gin"
)

func UserTypeRoutes(handler *handlers.UserTypeHandler, db *sql.DB, router *gin.RouterGroup) {
	router.GET("/user-types", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "User Type List"})
	})

	router.POST("/user-type", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Create User Type"})
	})

	router.GET("/user-types/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"message": "Return User Type", "id": id})
	})
}
