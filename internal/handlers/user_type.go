package handlers

import "github.com/eralves01/user-services/internal/services"

type UserTypeHandler struct {
	service *services.UserTypeService
}

func NewUserTypeHandler(service *services.UserTypeService) *UserTypeHandler {
	return &UserTypeHandler{service: service}
}
