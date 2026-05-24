package user

import (
	"database/sql"
	//"github.com/kholeur9/dhaclub-app/internal/apperrors"
	///"fmt"
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
	var userCreated User
	row := pu.db.QueryRow(`INSERT INTO users(id, email, username, password) VALUES($1, $2, $3, $4) RETURNING id, email, username, created_at, updated_at`, user.ID, user.Email, user.Username, user.Password)
	err := row.Scan(&userCreated.ID, &userCreated.Email, &userCreated.Username, &userCreated.CreatedAt, &userCreated.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &userCreated, nil
}

func (pu *PostgresUser) UserExistsByEmail(email string) (bool, error) {
	var exists string
	row := pu.db.QueryRow(`SELECT id FROM users WHERE email = $1`, email)
	err := row.Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
