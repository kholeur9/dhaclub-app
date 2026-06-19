package user

import (
	"encoding/json"
	"fmt"
	//"context"
	"net/http"

	"github.com/kholeur9/dhaclub-app/internal/shared"
	"github.com/kholeur9/dhaclub-app/internal/response"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service}
} 

func (uh *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	create := CreateUserDto{}
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

func (uh *UserHandler) UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var updatePasswordDto UpdatePasswordDto
	if err := json.NewDecoder(r.Body).Decode(&updatePasswordDto); err != nil {
		response.HandleServiceError(w, err)
		return
	}
	passwordUpdated, err := uh.service.UserUpdatePassword(id, updatePasswordDto)
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 200, passwordUpdated)
}

func (uh *UserHandler) GetMeHandler(w http.ResponseWriter, r *http.Request) {
	userIDContext := r.Context().Value(shared.UserIDKey)
	fmt.Println("Handler:", userIDContext)
}