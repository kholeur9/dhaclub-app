package apperrors

type ErrorType string

const (
	VALIDATION ErrorType = "VALIDATION"
	CONFLICT ErrorType = "CONFLICT"
	INTERNAL ErrorType = "INTERNAL"
	NOT_FOUND ErrorType = "NOT_FOUND"
	UNAUTHORIZED ErrorType = "UNAUTHORIZED"
	FORBIDDEN ErrorType = "FORBIDDEN"
)