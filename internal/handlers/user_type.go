package handlers

import (
	"net/http"
	"strconv"

	"github.com/eralves01/user-services/internal/domain"
	"github.com/eralves01/user-services/internal/services"
	"github.com/gin-gonic/gin"
)

type UserTypeHandler struct {
	service *services.UserTypeService
}

func NewUserTypeHandler(service *services.UserTypeService) *UserTypeHandler {
	return &UserTypeHandler{service: service}
}

func (h *UserTypeHandler) Create(c *gin.Context) {
	var userType domain.UserType

	if err := c.ShouldBindJSON(&userType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(userType); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User Type created successfully!"})
}

func (h *UserTypeHandler) GetByValue(c *gin.Context) {
	value := c.Query("value")
	if value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "value parameter is missing"})
		return
	}

	userType, err := h.service.GetByValue(value)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userType": userType})
}

func (h *UserTypeHandler) GetByid(c *gin.Context) {
	id, exists := c.Params.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is missing"})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter must be an integer"})
		return
	}

	userType, err := h.service.GetByID(idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userType": userType})
}

func (h *UserTypeHandler) GetAll(c *gin.Context) {
	userTypes, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userTypes": userTypes})
}
