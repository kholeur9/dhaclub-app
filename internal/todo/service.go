package todo

import (
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

func (ts *TodoService) CreateTodo(t CreateTodoDto) (*CreateTodoResponse, error) {
	// Verify if description is not registered
	description := strings.TrimSpace(t.Description)
	if description == "" {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: "Description is required.",
		}
	}
	// Verify the length
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
			Message: "Internal server error.",
		}
	}
	if todoExists {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.CONFLICT,
			Message: "Todo already exists.",
		}
	}
	createdID := uuid.New().String()
	newTodo := Todo{
		ID:          createdID,
		Description: description,
	}
	todo, err := ts.store.Add(newTodo)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: "Internal server error.",
		}
	}
	return &CreateTodoResponse{
		Message: "Succesfully",
		Data: TodoDto{
			ID:          todo.ID,
			Description: todo.Description,
			Done:        todo.IsDone,
		},
	}, nil
}

func (ts *TodoService) GetTodoByID(id string) (*Todo, error) {
	todo, err := ts.store.GetByID(id)
	if errors.Is(err, apperrors.ErrTodoNotFound) {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.NOT_FOUND,
			Message: "Todo not found.",
		}
	}
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: "Internal server error.",
		}
	}
	return todo, nil
}

func (ts *TodoService) TodosList() ([]*Todo, error) {
	getAllTodos, err := ts.store.TodosList()
	if err != nil {
		return nil, err
	}
	return getAllTodos, nil
}

func (ts *TodoService) DeleteTodo(id string) (*DeleteTodoResponse, error) {
	todoID, err := ts.store.DeleteTodo(id)
	if errors.Is(err, apperrors.ErrTodoNotFound) {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.NOT_FOUND,
			Message: "Todo not found.",
		}
	} 
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: "Internal server error.",
		}
	}
	return &DeleteTodoResponse{
		Message: "Succesfully",
		ID:      todoID,
	}, nil
}

func (ts *TodoService) UpdateTodo(id string, t UpdateTodoDto) (*Todo, error) {
	var updateField UpdateFieldDto
	if t.Description == nil && t.IsDone == nil {
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
		if len(*t.Description) <= 2 {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.VALIDATION,
				Message: "Description too short.",
			}
		}
	}
	todoExists, err := ts.store.GetByID(id)
	if errors.Is(err, apperrors.ErrTodoNotFound) {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.NOT_FOUND,
			Message: "Todo not found.",
		}
	}
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: "Internal server error.",
		}
	}
	if t.Description != nil && t.IsDone != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.CONFLICT,
			Message: "It is impossible to edit a todo and mark it as already done at the same time.",
		}
	}
	if t.IsDone != nil {
		isDone := *t.IsDone
		if isDone == todoExists.IsDone {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: "Nothing has changed.",
			}
		}
		updateField = UpdateFieldDto{
			ID:        id,
			IsDone:    &isDone,
			UpdatedAt: time.Now(),
		}
	}
	if t.Description != nil {
		description := strings.TrimSpace(*t.Description)
		if desc, err := ts.store.ExistsByDescription(description); desc {
			if err != nil {
				return nil, &apperrors.ServiceError{
					Type:    apperrors.INTERNAL,
					Message: "Internal server error.",
				}
			}
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: "This description is already a todo.",
			}
		}
		if description == todoExists.Description {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: "Nothing has changed.",
			}
		}
		updateField = UpdateFieldDto{
			ID:          id,
			Description: &description,
			UpdatedAt:   time.Now(),
		}
	}
	todoUpdated, err := ts.store.UpdateTodo(updateField)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: "Internal server error.",
		}
	}
	return todoUpdated, nil
}
