package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type RegisterRoutes struct {
	db     *sql.DB
	router *gin.Engine
}

func NewRegisterRoutes(db *sql.DB, router *gin.Engine) *RegisterRoutes {
	return &RegisterRoutes{db: db, router: router}
}

func (r *RegisterRoutes) GetRegisterRoutes() {
	api := r.router.Group("/api/v1")

	UserRoutes(r.db, api)
}
