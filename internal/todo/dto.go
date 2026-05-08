package todo

type TodoDto struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
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
	Message string `json:"message"`
	ID      *string `json:"id"`
}

type FieldModified struct {
	Description string `json:"description"`
	IsDone bool `json:"is_done"`
}
type UpdateTodoDto struct {
	ID string `json:"id"`
	Field FieldModified `json:"field"`
}