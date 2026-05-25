package user

import (
	"encoding/json"
	//"fmt"
	"net/http"

	"github.com/kholeur9/dhaclub-app/internal/response"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service}
} 

func (uh *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var create CreateUserDto
	if err := json.NewDecoder(r.Body).Decode(&create); err != nil {
		response.HandleServiceError(w, err)
		return
	}
	user, err := uh.service.CreateUser(create)
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 201, user)
}

func (uh *UserHandler) UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var update UpdateDto
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		response.HandleServiceError(w, err)
		return
	}
	userUpdated, err := uh.service.UserUpdate(id, update)
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 200, userUpdated)
}