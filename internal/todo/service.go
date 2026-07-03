package todo

import (
	//"fmt"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kholeur9/dhaclub-app/internal/apperrors"
)

type TodoService struct {
	store TodoStore
}

func NewTodoService(store TodoStore) *TodoService {
	return &TodoService{
		store,
	}
}

func (ts *TodoService) CreateTodo(userID uuid.UUID, t CreateTodoDto) (*CreateTodoResponse, error) {
	description := strings.TrimSpace(t.Description)
	if description == "" {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: "Description is required.",
		}
	}
	if len(description) <= 2 {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: "Description too short.",
		}
	}
	todoExists, err := ts.store.ExistsByDescription(description)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	if todoExists {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.CONFLICT,
			Message: "Todo already exists.",
		}
	}
	createdID := uuid.New()
	newTodo := Todo{
		ID:          createdID,
		UserID:      userID,
		Description: description,
	}
	todo, err := ts.store.Add(newTodo)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	return &CreateTodoResponse{
		Message: "Successfully",
		Data: TodoDto{
			ID:          todo.ID.String(),
			UserID:      todo.UserID.String(),
			Description: todo.Description,
			Completed:   todo.Completed,
		},
	}, nil
}

func (ts *TodoService) GetTodoByID(userID uuid.UUID, todoID string) (*Todo, error) {
	todoIDParse, err := uuid.Parse(todoID)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.NOT_FOUND,
			Message: apperrors.TodoNotIdentified,
		}
	}
	todo, err := ts.store.GetUserTodoByID(userID, todoIDParse)
	if errors.Is(err, apperrors.ErrTodoNotFound) {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.NOT_FOUND,
			Message: apperrors.ErrTodoNotFoundMessage,
		}
	}
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	return todo, nil
}

func (ts *TodoService) TodosList(userID uuid.UUID, page int, limit int, todoFilter TodoFilter, todoSort TodoSort) ([]*Todo, error) {
	// Pagination Logic
	if page < 1 {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: "Inavlid value page.",
		}
	}
	if limit < 1 {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: "Invalid value limit.",
		}
	}
	if limit > 20 {
		limit = 20
	}
	offset := (page - 1) * limit
	// tri logic
	if todoSort.Sort != nil && todoSort.Order == nil || todoSort.Sort == nil && todoSort.Order != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: "Sort and order are necessary..",
		}
	}
	if todoSort.Sort != nil && todoSort.Order != nil {
		sort := *todoSort.Sort
		if sort != "created_at" && sort != "updated_at" {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.VALIDATION,
				Message: "Invalid value sort.",
			}
		}
		order := *todoSort.Order
		if order != "desc" && order != "asc" {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.VALIDATION,
				Message: "Invalid value order.",
			}
		}
	}
	getAllTodos, err := ts.store.TodosList(userID, limit, offset, todoFilter, todoSort)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	return getAllTodos, nil
}

func (ts *TodoService) DeleteTodo(userID uuid.UUID, todoID string) (*DeleteTodoResponse, error) {
	todoIDParse, err := uuid.Parse(todoID)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.NOT_FOUND,
			Message: apperrors.TodoNotIdentified,
		}
	}
	err = ts.store.DeleteUserTodo(userID, todoIDParse)
	if errors.Is(err, apperrors.ErrTodoNotFound) {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.NOT_FOUND,
			Message: apperrors.ErrTodoNotFoundMessage,
		}
	}
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	return &DeleteTodoResponse{
		Message: "Todo deleted successfully.",
	}, nil
}

func (ts *TodoService) UpdateTodo(userID uuid.UUID, todoID string, t UpdateTodoDto) (*Todo, error) {
	var updateField UpdateFieldDto
	if t.Description == nil && t.Completed == nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: "No field to update.",
		}
	}
	if t.Description != nil {
		description := strings.TrimSpace(*t.Description)
		if description == "" {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.VALIDATION,
				Message: "Description is required.",
			}
		}
		if len(description) <= 2 {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.VALIDATION,
				Message: "Description too short.",
			}
		}
	}
	parsedTodoID, err := uuid.Parse(todoID)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.NOT_FOUND,
			Message: apperrors.TodoNotIdentified,
		}
	}
	todoExists, err := ts.store.GetUserTodoByID(userID, parsedTodoID)
	if errors.Is(err, apperrors.ErrTodoNotFound) {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.NOT_FOUND,
			Message: apperrors.ErrTodoNotFoundMessage,
		}
	}
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	if t.Description != nil && t.Completed != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.CONFLICT,
			Message: "It is impossible to edit a todo and mark it as already done at the same time.",
		}
	}
	if t.Completed != nil {
		completed := *t.Completed
		if completed == todoExists.Completed {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: "Nothing has changed.",
			}
		}
		updateField = UpdateFieldDto{
			ID:        parsedTodoID,
			Completed: &completed,
			UpdatedAt: time.Now(),
		}
	} else if t.Description != nil {
		description := strings.TrimSpace(*t.Description)
		if description == todoExists.Description {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: "Nothing has changed.",
			}
		}
		updateField = UpdateFieldDto{
			ID:          parsedTodoID,
			Description: &description,
			UpdatedAt:   time.Now(),
		}
	}
	todoUpdated, err := ts.store.UpdateTodo(userID, updateField)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	return todoUpdated, nil
}
