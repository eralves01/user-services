package domain

import "errors"

var (
	ErrUserTypeNotFound = errors.New("user type not found")
	ErrUserNotFound     = errors.New("user not found")
)
