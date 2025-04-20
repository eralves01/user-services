package domain

type UserType struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func NewUserType(id int, value string) *UserType {
	return &UserType{
		ID:    id,
		Value: value,
	}
}
