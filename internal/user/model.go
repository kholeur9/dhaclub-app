package user

import "time"

type User struct {
	ID           string     `json:"id"`
	Email        string     `json:"email"`
	Username     string     `json:"username"`
	Password     string     `json:"password"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAtdAt *time.Time `json:"updated_at"`
}

type UserStore interface {
	Create(user User) (*User, error)
}