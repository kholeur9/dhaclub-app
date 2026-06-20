package todo

import (
	//"fmt"
	//"context"
	"encoding/json"

	//"strings"

	//"strconv"
	//"html"
	"net/http"

	"github.com/google/uuid"
	"github.com/kholeur9/dhaclub-app/internal/apperrors"
	"github.com/kholeur9/dhaclub-app/internal/response"
	"github.com/kholeur9/dhaclub-app/internal/shared"
)

type HandlerTodo struct {
	todoService *TodoService
}

func NewHandlerTodo(todoService *TodoService) *HandlerTodo {
	return &HandlerTodo{todoService}
}

func (s *HandlerTodo) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	userIDContext := r.Context().Value(shared.UserIDKey)
	if userIDContext == nil {
		response.HandleServiceError(w, &apperrors.ServiceError{
			Type: apperrors.UNAUTHORIZED,
			Message: apperrors.UserNotAuthticated,
		})
		return
	}
	userID, ok := userIDContext.(uuid.UUID)
	if !ok {
		response.HandleServiceError(w, &apperrors.ServiceError{
			Type: apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		})
		return
	}
	// read and matched elements to body in struct
	var structData CreateTodoDto
	if readJSON := json.NewDecoder(r.Body).Decode(&structData); readJSON != nil {
		response.HandleServiceError(w, readJSON)
		return
	}
	// Send data at service
	todo, err := s.todoService.CreateTodo(userID, structData)
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 201, todo)
}

func (s *HandlerTodo) GetTodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	urlTodo := r.PathValue("id")
	todo, err := s.todoService.GetTodoByID(urlTodo)
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 200, todo)
}

func (s *HandlerTodo) TodosListHandler(w http.ResponseWriter, r *http.Request) {
	getAllTodos, err := s.todoService.TodosList()
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 200, getAllTodos)
}

func (s *HandlerTodo) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	result, err := s.todoService.DeleteTodo(id)
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 200, result)
}

func (s *HandlerTodo) UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var structData UpdateTodoDto
	if err := json.NewDecoder(r.Body).Decode(&structData); err != nil {
		response.HandleServiceError(w, err)
		return
	}
	//fmt.Printf("%+v\n", structData)
	result, err := s.todoService.UpdateTodo(id, structData)
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 200, result)
}