package domain

type UserType struct {
	ID    int
	Value string
}

func NewUserType(id int, value string) *UserType {
	return &UserType{
		ID:    id,
		Value: value,
	}
}
