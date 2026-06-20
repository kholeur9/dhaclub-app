package jwt

import (
	"time"

	//"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kholeur9/dhaclub-app/internal/apperrors"
)

type JwtService struct {
	key string
}

func NewJwtService(key string) *JwtService {
	return &JwtService{
		key: key,
	}
}

func (jws *JwtService) GenerateToken(username string, userID uuid.UUID) (string, error) {
	SigningKey := []byte(jws.key)
	claims := Claims{
		username,
		jwt.RegisteredClaims{
			Subject: userID.String(),
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

func (jws *JwtService) ValidateToken(token string) (*Claims, error) {
	tokenVerified, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte(jws.key), nil
	})
	if err != nil {
		return nil, err
	}
	if !tokenVerified.Valid {
		return nil, apperrors.ErrInvalidToken
	}
	if claims, ok := tokenVerified.Claims.(*Claims); ok {
		return claims, nil
	} else {
		return nil, apperrors.ErrClaimsUnknown
	}
}