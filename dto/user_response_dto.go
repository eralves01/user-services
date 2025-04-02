package dto

type UserResponseDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	UserTypeID int    `json:"user_type_id"`
}
