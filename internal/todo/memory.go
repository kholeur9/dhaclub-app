package todo

import (
	"errors"
	//"fmt"
)

type MemoryTodo struct {
	todos map[string]Todo
}

func NewMemoryTodo() *MemoryTodo {
	return &MemoryTodo{
		todos: map[string]Todo{},
	}
}

func (t *MemoryTodo) Add(td Todo) error {
	if _, exists := t.todos[td.ID]; exists {
		return errors.New("Todo exists")
	}
	t.todos[td.ID] = td
	return nil
}

func (t *MemoryTodo) GetByID(id string) (*Todo, error) {
	var todoResult *Todo
	for _, todo := range t.todos {
		if todo.ID == id {
			todoResult = &todo
		}
	}
	return todoResult, nil
}

func (t *MemoryTodo) TodosList() []Todo {
	var todosResult []Todo
	for _, todo := range t.todos {
		todosResult = append(todosResult, todo)
	}
	return todosResult
}

func (t *MemoryTodo) ExistsByDescription(desc string) (bool, error) {
	for _, todo := range t.todos {
		if todo.Description == desc {
			return true, errors.New("todo already exists")
		}
	}
	return false, nil
}