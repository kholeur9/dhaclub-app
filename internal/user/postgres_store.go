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
	var userCreated User
	row := pu.db.QueryRow(`INSERT INTO users(id, email, username, password) VALUES($1, $2, $3, $4) RETURNING id, email, username, created_at`, user.ID, user.Email, user.Username, user.Password)
	err := row.Scan(&userCreated.ID, &userCreated.Email, &userCreated.Username, &userCreated.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &userCreated, nil
}