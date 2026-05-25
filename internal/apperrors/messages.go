package apperrors

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