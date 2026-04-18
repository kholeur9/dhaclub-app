package todo

import (
	"errors"
	"time"
)

var (
	HttpNoValid = errors.New("this method is not allowed")
)

type Todo struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	IsDone      bool       `json:"done"`
	CreatedAt   time.Time  `json:"createdat"`
	UpdatedAt   *time.Time `json:"updatedat"`
}

type TodoStore interface {
	Add(t Todo) error
	ExistsByDescription(desc string) (bool, error)
	TodosList() []Todo
	GetByID(id string) (Todo, error)
}
