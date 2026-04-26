package todo

import "errors"

var (
	ErrTodoExists           = errors.New("todo already exists")
	ErrTodoNotFound         = errors.New("todo not found")
	ErrTodoTooShort         = errors.New("description too short")
	ErrDescriptionTodoEmpty = errors.New("must have an description")
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
	return se.Error()
}
