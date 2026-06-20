package todo

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID  `json:"id"`
	UserID      uuid.UUID  `json:"user_id"`
	Description string     `json:"description"`
	Completed   bool       `json:"is_done"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type TodoStore interface {
	Add(t Todo) (*Todo, error)
	ExistsByDescription(desc string) (bool, error)
	TodosList() ([]*Todo, error)
	GetByID(userID uuid.UUID, todoID string) (*Todo, error)
	DeleteTodo(id string) (*string, error)
	UpdateTodo(t UpdateFieldDto) (*Todo, error)
}
