package todo

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID  `json:"id"`
	UserID      uuid.UUID  `json:"user_id"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type TodoStore interface {
	Add(t Todo) (*Todo, error)
	ExistsByDescription(desc string) (bool, error)
	TodosList(userID uuid.UUID, limit int, offset int, todoFilter TodoFilter, todoTri TodoTri) ([]*Todo, error)
	GetUserTodoByID(userID, todoID uuid.UUID) (*Todo, error)
	DeleteUserTodo(userID uuid.UUID, todoID uuid.UUID) error
	UpdateTodo(userID uuid.UUID, t UpdateFieldDto) (*Todo, error)
}
