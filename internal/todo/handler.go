package todo

import (
	//"fmt"
	"strconv"
	//"strings"

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
			Type:    apperrors.UNAUTHORIZED,
			Message: apperrors.UserNotAuthticated,
		})
		return
	}
	userID, ok := userIDContext.(uuid.UUID)
	if !ok {
		response.HandleServiceError(w, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
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
	userIDContext := r.Context().Value(shared.UserIDKey)
	if userIDContext == nil {
		response.HandleServiceError(w, &apperrors.ServiceError{
			Type:    apperrors.UNAUTHORIZED,
			Message: apperrors.UserNotAuthticated,
		})
		return
	}
	userID, ok := userIDContext.(uuid.UUID)
	if !ok {
		response.HandleServiceError(w, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		})
		return
	}
	urlTodoID := r.PathValue("id")
	todo, err := s.todoService.GetTodoByID(userID, urlTodoID)
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 200, todo)
}

type TodoFilter struct {
	Completed *bool
}

type TodoSort struct {
	Sort *string
	Order *string
}

func (s *HandlerTodo) TodosListHandler(w http.ResponseWriter, r *http.Request) {
	var todoFilter TodoFilter
	userIDContext := r.Context().Value(shared.UserIDKey)
	if userIDContext == nil {
		response.HandleServiceError(w, &apperrors.ServiceError{
			Type:    apperrors.UNAUTHORIZED,
			Message: apperrors.UserNotAuthticated,
		})
		return
	}
	userID, ok := userIDContext.(uuid.UUID)
	if !ok {
		response.HandleServiceError(w, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		})
		return
	}
	// Pagination
	getURL := r.URL.Query()
	if !getURL.Has("page") {
		getURL.Add("page", "1")
	}
	if !getURL.Has("limit") {
		getURL.Add("limit", "20")
	}
	if getURL.Get("page") == "" {
		getURL.Set("page", "1")
	}
	if getURL.Get("limit") == "" {
		getURL.Set("limit", "20")
	}
	page, err := strconv.Atoi(getURL.Get("page"))
	if err != nil {
		response.HandleServiceError(w, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: "Invalid page number.",
		})
		return
	}
	limit, err := strconv.Atoi(getURL.Get("limit"))
	if err != nil {
		response.HandleServiceError(w, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: "Invalid limit value.",
		})
		return
	}
	// Filtered by completed or not completed
	if getURL.Has("completed") {
		value, err := strconv.ParseBool(getURL.Get("completed"))
		if err != nil {
			response.HandleServiceError(w, &apperrors.ServiceError{
				Type:    apperrors.VALIDATION,
				Message: "Invalid completed value.",
			})
			return
		}
		todoFilter = TodoFilter{
			Completed: &value,
		}
	}
	// Tried by sort : Created at, updated at And Order : DESC, ASC
	var todoSort TodoSort
	if getURL.Has("sort") || getURL.Has("order") {
		valueSort := getURL.Get("sort")
		valueOrder := getURL.Get("order")
		todoSort = TodoSort{
			Sort: &valueSort,
			Order: &valueOrder,
		}
	}
	getAllTodos, err := s.todoService.TodosList(userID, page, limit, todoFilter, todoSort)
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 200, getAllTodos)
}

func (s *HandlerTodo) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	userIDContext := r.Context().Value(shared.UserIDKey)
	if userIDContext == nil {
		response.HandleServiceError(w, &apperrors.ServiceError{
			Type:    apperrors.UNAUTHORIZED,
			Message: apperrors.UserNotAuthticated,
		})
		return
	}
	userID, ok := userIDContext.(uuid.UUID)
	if !ok {
		response.HandleServiceError(w, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		})
		return
	}
	todoID := r.PathValue("id")
	result, err := s.todoService.DeleteTodo(userID, todoID)
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 200, result)
}

func (s *HandlerTodo) UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	userIDContext := r.Context().Value(shared.UserIDKey)
	if userIDContext == nil {
		response.HandleServiceError(w, &apperrors.ServiceError{
			Type:    apperrors.UNAUTHORIZED,
			Message: apperrors.UserNotAuthticated,
		})
		return
	}
	userID, ok := userIDContext.(uuid.UUID)
	if !ok {
		response.HandleServiceError(w, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		})
		return
	}
	todoID := r.PathValue("id")
	var structData UpdateTodoDto
	if err := json.NewDecoder(r.Body).Decode(&structData); err != nil {
		response.HandleServiceError(w, err)
		return
	}
	//fmt.Printf("%+v\n", structData)
	result, err := s.todoService.UpdateTodo(userID, todoID, structData)
	if err != nil {
		response.HandleServiceError(w, err)
		return
	}
	response.WriteResponse(w, 200, result)
}
