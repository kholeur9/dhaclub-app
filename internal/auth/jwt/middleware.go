package jwt

import (
	"context"
	"fmt"
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
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}
		verify := strings.Split(token, " ")
		fmt.Println(verify)
		if len(verify) < 1 {
			http.Error(w, "Empty authorization", http.StatusUnauthorized)
			return
		}
		if verify[0] != "Bearer" {
			http.Error(w, "Missing bearer", http.StatusUnauthorized)
			return
		}
		tokenHeader := verify[1]
		claims, err := ams.jwtService.ValidateToken(tokenHeader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(context.Background(), shared.UserIDKey, claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}