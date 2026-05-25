package user

import (
	"database/sql"
	"errors"

	"github.com/kholeur9/dhaclub-app/internal/apperrors"
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

func (pu *PostgresUser) FindConflicts(email, username *string) (*User, error) {
	var user User
	row := pu.db.QueryRow(`SELECT id, email, username FROM users WHERE email = $1 OR username = $2`, email, username)
	err := row.Scan(&user.ID, &user.Email, &user.Username)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (pu *PostgresUser) UserExistsByEmail(email string) (bool, error) {
	var id string
	row := pu.db.QueryRow(`SELECT id FROM users WHERE email = $1`, email)
	err := row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (pu *PostgresUser) GetUserById(id string) (*GetUserResponseDto, error) {
	var user GetUserResponseDto
	row := pu.db.QueryRow(`SELECT id, email, username, created_at, updated_at FROM users WHERE id = $1`, id)
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, apperrors.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (pu *PostgresUser) Update(dto UpdateResponse) (*User, error) {
	var user User
	// I use COALESCE because i want to change value not null
	row := pu.db.QueryRow(`UPDATE users SET COALESCE($1, email), COALESCE($2, username), updated_at = $3 WHERE id = $4 RETURNING id, email, updated_at`, dto.Email, dto.Username, dto.UpdatedAt, dto.ID)
	err := row.Scan(&user.ID, &user.Email, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}