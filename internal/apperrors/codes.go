package apperrors

import "errors"

// USER
var (
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