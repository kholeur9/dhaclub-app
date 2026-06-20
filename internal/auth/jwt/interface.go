package jwt

import "github.com/google/uuid"

type Token string

type JWTService interface {
	GenerateToken(username string, userID uuid.UUID) (string, error)
	ValidateToken(token string) (*Claims, error)
}