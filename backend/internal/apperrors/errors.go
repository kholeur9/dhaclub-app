package apperrors

type ServiceError struct {
	Type ErrorType
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}