package todo

import (
	//"fmt"
	//"context"
	"encoding/json"

	//"strconv"
	//"html"
	"net/http"

	"github.com/kholeur9/dhaclub-app/internal/response"
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
		response.HandleServiceError(w, readJSON)
		return
	}
	// Send data at service
	todo, err := s.todoService.CreateTodo(structData)
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