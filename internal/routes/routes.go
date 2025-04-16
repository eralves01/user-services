package routes

import (
	"database/sql"

	"github.com/eralves01/user-services/internal/handlers"
	"github.com/eralves01/user-services/internal/repositories"
	"github.com/eralves01/user-services/internal/services"
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
	userTypeRepository := repositories.NewUserTypeRepository(r.db)
	userTypeService := services.NewUserTypeService(userTypeRepository)
	userTypeHandler := handlers.NewUserTypeHandler(userTypeService)

	userRepository := repositories.NewUserRepository(r.db)
	userService := services.NewUserService(userRepository, userTypeService)
	userHandler := handlers.NewUserHandler(userService)

	api := r.router.Group("/api/v1")

	UserRoutes(userHandler, r.db, api)
	UserTypeRoutes(userTypeHandler, r.db, api)
}
