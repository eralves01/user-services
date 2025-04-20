package dto

import (
	"errors"
	"regexp"
	"strings"

	"github.com/eralves01/user-services/internal/domain"
)

type CreateUserInput struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	UserTypeID int    `json:"user_type_id"`
}

func (c *CreateUserInput) Validate() error {
	var errorsList []string

	// Validações
	if len(c.Name) < 3 {
		errorsList = append(errorsList, "Name must be at least 3 characters long")
	}
	if !isValidEmail(c.Email) {
		errorsList = append(errorsList, "Invalid email format")
	}
	if len(c.Password) < 6 {
		errorsList = append(errorsList, "Password must be at least 6 characters long")
	}
	// Se houver erros, retorna todos concatenados
	if len(errorsList) > 0 {
		return errors.New(strings.Join(errorsList, "; "))
	}

	return nil
}

func (c *CreateUserInput) ToUser() *domain.User {
	return domain.NewUser(c.Name, c.Email, c.UserTypeID, c.Password)
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
