package todo

import (
	//"context"
	"dhaclub-app/internal/todo/utils"
	"encoding/json"

	//"strconv"
	//"html"
	"net/http"
)

type HandlerTodo struct {
	todoService *TodoService
}

func NewHandlerTodo(todoService *TodoService) *HandlerTodo {
	return &HandlerTodo{todoService}
}

func (s *HandlerTodo) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	// taked json of body
	clientData := r.Body
	// read and matched elements to body in struct
	structData := CreateTodoDto{}
	if readJSON := json.NewDecoder(clientData).Decode(&structData); readJSON != nil {
		http.Error(w, "JSON failed", http.StatusBadRequest)
		return
	}
	// Send data at service
	todo, err := s.todoService.CreateTodo(structData)
	if err != nil {
		if v, ok := err.(*ServiceError); ok {
			switch v.Type {
			case "VALIDATION":
				utils.WriteError(w, v.Message, http.StatusBadRequest)
				return
			case "CONFLICT":
				utils.WriteError(w, v.Message, http.StatusConflict)
				return
			default:
				utils.WriteError(w, v.Message, http.StatusInternalServerError)
				return
			}
		} else {
			utils.WriteError(w, "une erreur est survenue", http.StatusInternalServerError)
			return
		}
	}
	utils.WriteResponse(w, 201, todo)
}

func (s *HandlerTodo) GetTodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	urlTodo := r.PathValue("todo_id")
	todo, err := s.todoService.GetTodoByID(urlTodo)
	if err != nil {
		if val, ok := err.(*ServiceError); ok {
			switch val.Type {
			case "NOT_FOUND":
				utils.WriteError(w, val.Message, http.StatusNotFound)
				return
			}
		}
	}
	utils.WriteResponse(w, 200, todo)
}

func (s *HandlerTodo) TodosListHandler(w http.ResponseWriter, r *http.Request) {
	getAllTodos := s.todoService.TodosList()
	utils.WriteResponse(w, 200, getAllTodos)
}
