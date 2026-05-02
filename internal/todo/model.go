package todo

import (
	"time"
)

type Todo struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	IsDone      bool       `json:"is_done"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type TodoStore interface {
	Add(t Todo) (*Todo, error)
	ExistsByDescription(desc string) (bool, error)
	TodosList() ([]*Todo, error)
	GetByID(id string) (*Todo, error)
}
