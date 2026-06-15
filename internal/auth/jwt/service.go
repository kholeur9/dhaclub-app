package jwt

import (
	"time"
	//"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	key string
}

func NewJwtService(key string) *jwtService {
	return &jwtService{
		key: key,
	}
}

func (jws *jwtService) GenerateToken(userID string) (string, error) {
	SigningKey := jws.key
	claims := Claims{
		userID,
		jwt.RegisteredClaims{
			Subject: userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	tokenGenerated := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenGenerated.SignedString(SigningKey)
	if err != nil {
		return "", err
	}
	return token, nil
}