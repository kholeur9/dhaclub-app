package jwt

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Username string `json:"user_id"`
	jwt.RegisteredClaims
}