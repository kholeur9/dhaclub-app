package todo

import (
	"errors"
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
	if t.Description == "" {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: "Must have an description",
		}
	}
	// Verify the length
	if len(t.Description) <= 2 {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: "Description too short",
		}
	}
	todoExists, err := ts.store.ExistsByDescription(t.Description)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: "Internal server error",
		}
	}
	if todoExists {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.CONFLICT,
			Message: "Todo already exists",
		}
	}
	createdID := uuid.New().String()
	newTodo := Todo{
		ID:          createdID,
		Description: t.Description,
	}
	todo, err := ts.store.Add(newTodo)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: "Internal server error",
		}
	}
	return &CreateTodoResponse{
		Message: "Création réussie",
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
			Message: "Todo not found",
		}
	} else if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: "Internal server error",
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
			Message: "todo not found.",
		}
	} else if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: "Internal server error",
		}
	}
	return &DeleteTodoResponse{
		Message: "Todo deleted succesfuly",
		ID:      todoID,
	}, nil
}

func (ts *TodoService) UpdateTodo(id string, t UpdateTodoDto) (*Todo, error) {
	var updateField UpdateFieldDto
	todoExists, err := ts.store.GetByID(id)
	if errors.Is(err, apperrors.ErrTodoNotFound) {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.NOT_FOUND,
			Message: "Todo does not exist.",
		}
	} else if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: "Internal server error",
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
	} else if t.Description != nil {
		description := *t.Description
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
	return todoUpdated, nil
}
