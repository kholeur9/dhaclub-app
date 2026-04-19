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

