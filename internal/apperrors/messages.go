package apperrors

// GENERIC
const ErrInternalServerErrorMessage = "Internal server error"

// USER
const (
	UserNotFoundMessage   = "This account does not exists."
	IdentifierRequired    = "Identifier is required."
	EmailRequired         = "Email is required."
	UsernameRequired      = "Username is required."
	PasswordRequired      = "Password is required."
	UsernameShort         = "Username too short."
	PasswordShort         = "Password must be at least 8 characters."
	EmailAlreadyExists    = "Email already exists."
	UsernameAlreadyExists = "Username already exists."
	EmailNotChanged       = "You do not updated your email."
	UsernameNotChanged    = "You do not updated your usernme."
	OneFieldToUpdate      = "You can only update one field at a time."
	NoDataToUpdate        = "No data to update."

	InvalidCredentials = "Identifier or password invalid."
	AccountNotActive   = "Account not active."

	UserNotAuthticated = "User not authenticated."

	AccessDenied = "You do not have permission to perform this action."
)

// TODO
const (
	ErrTodoNotFoundMessage = "Todo not found"
)
