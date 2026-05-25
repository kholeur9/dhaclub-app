package apperrors

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrManyFieldsToUpdate = errors.New("many fields to update")
	ErrNoFieldsToUpdate   = errors.New("no fields to update")
)