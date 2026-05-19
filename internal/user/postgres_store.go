package user

import (
	"database/sql"
	"fmt"
)

type PostgresUser struct {
	db *sql.DB
}

func NewPostgresUser(db *sql.DB) *PostgresUser {
	return &PostgresUser{
		db: db,
	}
}

func (pu *PostgresUser) Create(user User) (*User, error) {
	fmt.Println("Postgres user :", user)
	return nil, nil
}