package routes

import (
	"database/sql"

	"github.com/eralves01/user-services/internal/handlers"
	"github.com/gin-gonic/gin"
)

func UserTypeRoutes(handler *handlers.UserTypeHandler, db *sql.DB, router *gin.RouterGroup) {
	router.POST("/user-type", handler.Create)

	router.GET("/user-type/:id", handler.GetByid)

	router.GET("/user-type", handler.GetByValue)

	router.GET("/user-types", handler.GetAll)
}
