package todo

import (
	//"context"
	"encoding/json"
	//"fmt"

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
	w.Header().Set("Content-Type", "application/json")
	// Verify method request
	if r.Method != "POST" {
		http.Error(w, "this method is not allowed", http.StatusMethodNotAllowed)
		return
	}
	// taked json of body
	clientData := r.Body
	// read and matched elements to body in struct
	structData := CreateTodoDto{}
	if readJSON := json.NewDecoder(clientData).Decode(&structData); readJSON != nil {
		http.Error(w, "JSON failed", http.StatusBadRequest)
	}
	// Send data at service
	todo, err := s.todoService.CreateTodo(structData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&todo)
}

func (s *HandlerTodo) GetTodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, HttpNoValid.Error(), http.StatusMethodNotAllowed)
		return
	}
	urlTodo := r.PathValue("todo_id")
	todo, err := s.todoService.GetTodoByID(urlTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&todo)
}

func (s *HandlerTodo) TodosListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, HttpNoValid.Error(), http.StatusMethodNotAllowed)
		return
	}
	getAllTodos := s.todoService.TodosList()
	json.NewEncoder(w).Encode(&getAllTodos)
}