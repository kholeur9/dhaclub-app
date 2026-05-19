package helpers

import "golang.org/x/crypto/bcrypt"

type PasswordSecured interface {
	PasswordHash(password string) (string, error)
}

func PasswordHash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}