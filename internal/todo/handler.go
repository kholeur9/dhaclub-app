package todo

import (
	//"context"
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
		HandleServiceError(err)
	}
	WriteResponse(w, 201, todo)
}

func (s *HandlerTodo) GetTodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	urlTodo := r.PathValue("todo_id")
	todo, err := s.todoService.GetTodoByID(urlTodo)
	if err != nil {
		HandleServiceError(err)
	}
	WriteResponse(w, 200, todo)
}

func (s *HandlerTodo) TodosListHandler(w http.ResponseWriter, r *http.Request) {
	getAllTodos := s.todoService.TodosList()
	WriteResponse(w, 200, getAllTodos)
}
