package dto

import (
	"errors"
	"regexp"
	"strings"

	"github.com/eralves01/user-services/constants"
)

type CreateUserDTO struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordHash string
	UserType     constants.UserType `json:"user_type"`
	UserTypeID   int
}

func (c *CreateUserDTO) Validate() error {
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
	if !isValidUserType(c.UserType) {
		errorsList = append(errorsList, "The user type is not valid")
	}

	// Se houver erros, retorna todos concatenados
	if len(errorsList) > 0 {
		return errors.New(strings.Join(errorsList, "; "))
	}

	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidUserType(userType constants.UserType) bool {
	validTypes := map[constants.UserType]bool{
		constants.Client:   true,
		constants.Merchant: true,
	}
	return validTypes[userType]
}
