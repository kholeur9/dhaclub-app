package jwt

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/kholeur9/dhaclub-app/internal/shared"
)

type MiddlewareService struct {
	jwtService *JwtService
}

func NewAuthMiddlewareService(jwt *JwtService) *MiddlewareService{
	return &MiddlewareService{jwt}
}

func (ams *MiddlewareService) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 {
			http.Error(w, "Invalid format", http.StatusUnauthorized)
			return
		}
		if tokenString[0] != "Bearer" {
			http.Error(w, "Missing bearer", http.StatusUnauthorized)
			return
		}
		token := tokenString[1]
		if token == "" {
			http.Error(w, "Missing Token in authorization", http.StatusUnauthorized)
			return
		}
		claims, err := ams.jwtService.ValidateToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		subject, err := uuid.Parse(claims.Subject)
		if err != nil {
			http.Error(w, "Invalid user identity", http.StatusUnauthorized)
		}
		ctx := context.WithValue(r.Context(), shared.UserIDKey, subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}