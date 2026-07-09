package user

import "time"

type CreateUserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserResponseDto struct {
	Message string             `json:"message"`
	Data    CreateUserResponse `json:"data"`
}

type CreateUserDto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserDto struct {
	Identifer string `json:"identifier"`
	Password  string `json:"password"`
}

type LoginResponse struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type LoginUserResponse struct {
	Message     string        `json:"message"`
	AccessToken string        `json:"access_token"`
	Data        LoginResponse `json:"data"`
}

type GetUserResponseDto struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UpdateDto struct {
	Email    *string `json:"email"`
	Username *string `json:"username"`
}
type UpdateResponse struct {
	ID        string    `json:"id"`
	Email     *string   `json:"email"`
	Username  *string   `json:"username"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateResponseDto struct {
	Message string         `json:"message"`
	Data    UpdateResponse `json:"data"`
}

type UpdatePasswordDto struct {
	OldPassword *string `json:"old_password"`
	NewPassword string  `json:"new_password"`
}

type UpdatePasswordResponse struct {
	Message string `json:"message"`
}
