package todo

import (
	"errors"
	//"fmt"
	"time"

	"github.com/google/uuid"
)

type TodoService struct {
	store TodoStore
}

func NewTodoService(store TodoStore) *TodoService {
	return &TodoService{store}
}

func (ts *TodoService) CreateTodo(t CreateTodoDto) (*CreateTodoResponse, error) {
	// Verify if description is not registered
	if t.Description == "" {
		return nil, errors.New("todo would have an description")
	}
	// Verify the length
	if len(t.Description) < 2 {
		return nil, errors.New("todo would have more two caracters")
	}
	todoExists, err := ts.store.ExistsByDescription(t.Description)
	if todoExists {
		return nil, err
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
		return nil, err
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

func (ts *TodoService) GetTodoByID(id string) (Todo, error) {
	if id == "" {
		return Todo{}, errors.New("un id est nécessaire pour trouver votre todo")
	}
	thisID, err := ts.store.GetByID(id)
	if err != nil {
		return Todo{}, err
	}
	return thisID, nil
}

func (ts *TodoService) TodosList() []Todo {
	getAllTodos := ts.store.TodosList() 
	return getAllTodos
}