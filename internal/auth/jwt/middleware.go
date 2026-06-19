package jwt

import (
	"context"
	"net/http"
	"strings"

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
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}
		verify := strings.Split(tokenString, " ")
		if len(verify) != 2 {
			http.Error(w, "Invalid format", http.StatusUnauthorized)
			return
		}
		if verify[0] != "Bearer" {
			http.Error(w, "Missing bearer", http.StatusUnauthorized)
			return
		}
		tokenHeader := verify[1]
		if tokenHeader == "" {
			http.Error(w, "Missing Token in authorization", http.StatusUnauthorized)
			return
		}
		claims, err := ams.jwtService.ValidateToken(tokenHeader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), shared.UserIDKey, claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}