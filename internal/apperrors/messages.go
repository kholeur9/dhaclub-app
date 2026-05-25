package apperrors

// GENERIC
const ErrInternalServerErrorMessage = "Internal server error"

// USER
const (
	UserNotFoundMessage   = "This account does not exists"
	EmailRequired         = "Email is required"
	UsernameRequired      = "Username is required"
	PasswordRequired      = "Password is required"
	UsernameShort         = "Username too short"
	PasswordShort         = "Password must be at least 8 characters"
	EmailAlreadyExists    = "Email already exists"
	UsernameAlreadyExists = "Username already exists"
	EmailNotChanged       = "You do not updated your email"
	UsernameNotChanged    = "You do not updated your usernme"
	OneFieldToUpdate      = "You can only update one field at a time"
	NoDataToUpdate        = "No data to update"
)

// TODO
const (
	ErrTodoNotFoundMessage = "Todo not found"
)
