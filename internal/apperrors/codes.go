package apperrors

import "errors"

// USER
var (
	ErrIdentifierEmpty = errors.New("Identifer empty")
	ErrPasswordEmpty = errors.New("Password empty")
	ErrAccountNotActive = errors.New("Account not active")
	ErrPasswordWrong = errors.New("Password wrong")
	ErrUserNotFound = errors.New("User not found")
	ErrManyFieldsToUpdate = errors.New("Many fields found")
	ErrNoFieldsToUpdate = errors.New("No data found")

)
// JWT
var (
	ErrInvalidToken = errors.New("Invalid token")
	ErrClaimsUnknown = errors.New("Unknown cleams type")
)

// TODO
var (
	ErrTodoExists           = errors.New("todo already exists")
	ErrTodoNotFound         = errors.New("todo not found")
	ErrTodoTooShort         = errors.New("description too short")
	ErrDescriptionTodoEmpty = errors.New("description is empty")
)