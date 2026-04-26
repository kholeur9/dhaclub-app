package todo

import (
	//"fmt"
	"time"
	"errors"

	"github.com/google/uuid"
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
		return nil, &ServiceError{
			Type: VALIDATION,
			Message: "Must have an description",
		}
	}
	// Verify the length
	if len(t.Description) < 2 {
		return nil, &ServiceError{
			Type: VALIDATION,
			Message: "Description too short",
		}
	}
	todoExists, err := ts.store.ExistsByDescription(t.Description)
	if err != nil {
		return nil, &ServiceError{
			Type: INTERNAL,
			Message: "Internal server error",
		}
	}
	if todoExists {
		return nil, &ServiceError{
			Type: CONFLICT,
			Message: "Todo already exists",
		}
	}
	createdID := uuid.New().String()
	newTodo := Todo{
		ID: createdID,
		Description: t.Description,
		IsDone: false,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
	err = ts.store.Add(newTodo)
	if err != nil {
		return nil, &ServiceError{
			Type: INTERNAL,
			Message: "Internal server error",
		}
	}
	return &CreateTodoResponse{
		Message: "Création réussie",
		Data: TodoDto{
			ID: newTodo.ID,
			Description: newTodo.Description,
			Done: newTodo.IsDone,
		},
	}, nil
}

func (ts *TodoService) GetTodoByID(id string) (*Todo, error) {
	thisID, err := ts.store.GetByID(id)
	if errors.Is(err, ErrTodoNotFound) {
		return nil, &ServiceError{
			Type: NOT_FOUND,
			Message: "Todo not found",
		}
	}
	return thisID, nil
}

func (ts *TodoService) TodosList() []*Todo {
	getAllTodos := ts.store.TodosList() 
	return getAllTodos
}