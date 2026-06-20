package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserStore interface {
	Create(user User) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	FindConflicts(email, username *string) (*User, error)
	UserExistsByEmail(email string) (bool, error)
	GetUserById(id string) (*GetUserResponseDto, error)
	Update(dto UpdateResponse) (*User, error)
}
