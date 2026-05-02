package todo

import (

	"github.com/kholeur9/dhaclub-app/internal/apperrors"
)

//"fmt"

type MemoryTodo struct {
	todos map[string]*Todo
}

func NewMemoryTodo() *MemoryTodo {
	return &MemoryTodo{
		todos: map[string]*Todo{},
	}
}

func (t *MemoryTodo) Add(td Todo) error {
	if _, exists := t.todos[td.ID]; exists {
		return apperrors.ErrTodoExists
	}
	t.todos[td.ID] = &td
	return nil
}

func (t *MemoryTodo) GetByID(id string) (*Todo, error) {
	if _, exists := t.todos[id]; exists {
		return t.todos[id], nil
	}
	return nil, apperrors.ErrTodoNotFound
}

func (t *MemoryTodo) TodosList() []*Todo {
	var todosResult []*Todo
	for _, todo := range t.todos {
		todosResult = append(todosResult, todo)
	}
	return todosResult
}

func (t *MemoryTodo) ExistsByDescription(desc string) (bool, error) {
	for _, todo := range t.todos {
		if todo.Description == desc {
			return true, nil
		}
	}
	return false, nil
}
