package apperrors

import "errors"

// USER
var (
	ErrIdentifierEmpty = errors.New("identifer empty")
	ErrPasswordEmpty = errors.New("password empty")
	ErrAccountNotActive = errors.New("Account not active")
	ErrPasswordWrong = errors.New("password wrong")
	ErrUserNotFound = errors.New("user not found")
	ErrManyFieldsToUpdate = errors.New("many fields found")
	ErrNoFieldsToUpdate = errors.New("no data found")
)

// TODO
var (
	ErrTodoExists           = errors.New("todo already exists")
	ErrTodoNotFound         = errors.New("todo not found")
	ErrTodoTooShort         = errors.New("description too short")
	ErrDescriptionTodoEmpty = errors.New("description is empty")
)