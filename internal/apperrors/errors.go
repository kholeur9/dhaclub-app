package apperrors

import "errors"

/* ================= INTERNAL ERRORS ================= */

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrManyFieldsToUpdate = errors.New("many fields to update")
	ErrNoFieldsToUpdate   = errors.New("no fields to update")
)

/* ================= ERROR TYPES ================= */

const (
	VALIDATION = "VALIDATION"
	CONFLICT   = "CONFLICT"
	INTERNAL   = "INTERNAL"
	NOT_FOUND  = "NOT_FOUND"
)

/* ================= USER MESSAGES ================= */

const (
	ErrInternalServerErrorMessage = "Internal server error"
	ErrUserNotFoundMessage        = "User not found"

	EmailRequired    = "Email is required"
	UsernameRequired = "Username is required"
	UsernameShort    = "Username must be at least 2 characters"
	PasswordRequired = "Password is required"
	PasswordShort    = "Password must be at least 8 characters"

	EmailAlreadyExists    = "Email already exists"
	UsernameAlreadyExists = "Username already exists"

	EmailNotChanged    = "Email has not changed"
	UsernameNotChanged = "Username has not changed"

	OneFieldToUpdate = "You can only update one field at a time"
	NoDataToUpdate   = "No data to update"
)