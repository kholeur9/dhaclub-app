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
		UserID:		 userID,
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
			UserID:		 todo.UserID.String(),
			Description: todo.Description,
			Done:        todo.Completed,
		},
	}, nil
}

func (ts *TodoService) GetTodoByID(id string) (*Todo, error) {
	todo, err := ts.store.GetByID(id)
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

func (ts *TodoService) TodosList() ([]*Todo, error) {
	getAllTodos, err := ts.store.TodosList()
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	return getAllTodos, nil
}

func (ts *TodoService) DeleteTodo(id string) (*DeleteTodoResponse, error) {
	todoID, err := ts.store.DeleteTodo(id)
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
		Message: "Successfully",
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
		if len(description) <= 2 {
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
			Message: apperrors.ErrTodoNotFoundMessage,
		}
	}
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
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
		if isDone == todoExists.Completed {
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
	} else if t.Description != nil {
		description := strings.TrimSpace(*t.Description)
		if description == todoExists.Description {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: "Nothing has changed.",
			}
		}
		exists, err := ts.store.ExistsByDescription(description)
		if err != nil {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.INTERNAL,
				Message: apperrors.ErrInternalServerErrorMessage,
			}
		}
		if exists {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: "This description is already a todo.",
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
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	return todoUpdated, nil
}
