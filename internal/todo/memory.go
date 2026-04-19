package todo

import (
	//"fmt"
)

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
		return ErrTodoExists
	}
	t.todos[td.ID] = &td
	return nil
}

func (t *MemoryTodo) GetByID(id string) (*Todo, error) {
	if _, exists := t.todos[id]; exists {
		return t.todos[id], nil
	}
	return nil, ErrTodoNotFound
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
			return true, ErrTodoExists
		}
	}
	return false, nil
}