package todo

import "time"

type TodoDto struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
type CreateTodoDto struct {
	Description string `json:"description"`
}
type GetTodoByIdDto struct {
	ID string `json:"id"`
}
type CreateTodoResponse struct {
	Message string  `json:"message"`
	Data    TodoDto `json:"data"`
}
type DeleteTodoResponse struct {
	Message string  `json:"message"`
	ID      *string `json:"id"`
}
type UpdateTodoDto struct {
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}
type UpdateFieldDto struct {
	ID          string    `json:"id"`
	Description *string   `json:"description"`
	Completed   *bool     `json:"completed"`
	UpdatedAt   time.Time `json:"updated_at"`
}
