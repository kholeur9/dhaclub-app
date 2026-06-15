package jwt

type Token string

type JWTService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (*Claims, error)
}