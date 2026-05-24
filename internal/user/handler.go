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

func (hu *UserHandler) EmailUpdateHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var emailUpdate EmailUpdateDto
	if err := json.NewDecoder(r.Body).Decode(&emailUpdate); err != nil {
		response.HandleServiceError(w, err)
		return
	}
	fmt.Println("id handler:", id)
	fmt.Println("email handler:", emailUpdate)
	userUpdated, err := hu.service.UpdateUserEmail(id, emailUpdate)
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 200, userUpdated)
}