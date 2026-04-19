package todo

import "errors"

var (
	ErrTodoExists   = errors.New("todo already exists")
	ErrTodoNotFound = errors.New("todo not found")
	ErrTodoTooShort = errors.New("description to short")
	ErrDescriptionTodoEmpty = errors.New("must have an description")
)