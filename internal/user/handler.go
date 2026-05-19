package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kholeur9/dhaclub-app/internal/response"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(userService *UserService) *UserHandler {
	return &UserHandler{
		userService,
	}
}

func (hu *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var create CreateUserDto
	if err := json.NewDecoder(r.Body).Decode(&create); err != nil {
		response.HandleServiceError(w, err)
		return
	}
	fmt.Println("Handler create user :", create)
	user, err := hu.service.CreateUser(create)
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 201, user)
}