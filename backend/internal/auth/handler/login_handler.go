package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kholeur9/dhaclub-app/internal/auth/service"
	"github.com/kholeur9/dhaclub-app/internal/response"
	"github.com/kholeur9/dhaclub-app/internal/user"
)

type AuthHandler struct{
	authService *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (uh *AuthHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	login := user.LoginUserDto{}
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		response.HandleServiceError(w, err)
		return
	}
	fmt.Println("Login: ", login)
	user, err := uh.authService.LoginUser(login)
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 200, user)
}