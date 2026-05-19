package user

import "fmt"

type PostgresUser struct {}

func NewPostgresUser() *PostgresUser {
	return &PostgresUser{}
}

func (pu *PostgresUser) Create(user User) (*User, error) {
	fmt.Println("Postgres user :", user)
	return nil, nil
}