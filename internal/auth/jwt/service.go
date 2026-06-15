package jwt

import (
	"errors"
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
	SigningKey := []byte(jws.key)
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

func (jws *jwtService) ValidateToken(token string) (*Claims, error) {
	tokenVerified, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte(jws.key), nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := tokenVerified.Claims.(*Claims); ok {
		return claims, nil
	} else {
		return nil, errors.New("Claims unknown, do not proceed")
	}
}