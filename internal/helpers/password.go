package helpers

import "golang.org/x/crypto/bcrypt"

type PasswordSecure interface {
	PasswordHash(password string) (string, error)
}

type Bcrypt struct {
	bcryptSecure PasswordSecure
}

func NewBcryptSecure() *Bcrypt {
	return &Bcrypt{}
}

func (b *Bcrypt) PasswordHash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}