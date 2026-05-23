package apperrors

import "errors"

var (
	ErrTodoExists           = errors.New("todo already exists")
	ErrTodoNotFound         = errors.New("todo not found")
	ErrTodoTooShort         = errors.New("description too short")
	ErrDescriptionTodoEmpty = errors.New("must have an description")

	ErrUserNotFound = errors.New("User Not Found")
)

const (
	ErrInternalServerErrorMessage = "Internal server error."
	ErrTodoNotFoundMessage        = "Todo not found."

	EmailRequired = "Email is required"
	UsernameRequired = "Username is required"
	UsernameShort = "Username too short, enter more one caractere."
	PasswordRequired = "Password is required"
	PasswordShort = "Password would have 8 caracteres minimum."
	UserAlreadyExists = "User already exists"
)

const (
	VALIDATION = "VALIDATION"
	CONFLICT   = "CONFLICT"
	INTERNAL   = "INTERNAL"
	NOT_FOUND  = "NOT_FOUND"
)

type ServiceError struct {
	Type    string
	Message string
}

func (se *ServiceError) Error() string {
	return se.Message
}
