package todo

import (
	"database/sql"

	"github.com/kholeur9/dhaclub-app/internal/apperrors"
)

type PostgresTodo struct {
	db *sql.DB
}

func NewPostgresTodo(db *sql.DB) *PostgresTodo {
	return &PostgresTodo{
		db: db,
	}
}

func (pt *PostgresTodo) Add(t Todo) (*Todo, error) {
	var todoCreated *Todo
	row := pt.db.QueryRow(`INSERT INTO todos(id, description) VALUES($1, $2)`, t.ID, t.Description)
	if err := row.Scan(&todoCreated.ID, &todoCreated.Description, &todoCreated.IsDone); err != nil {
		return nil, err
	}
	return todoCreated, nil
}

func (pt *PostgresTodo) ExistsByDescription(description string) (bool, error) {
	var desc string
	todo := pt.db.QueryRow(`SELECT 1 FROM todos WHERE description = $1`, description)
	err := todo.Scan(&desc)
	if err == nil {
		return true, nil
	}
	if err == sql.ErrNoRows {
		return false, nil
	}
	return false, err
}

func (pt *PostgresTodo) TodosList() ([]*Todo, error) {
	rows, err := pt.db.Query(`SELECT id, description, is_done, created_at, updated_at FROM todos ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []*Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Description, &todo.IsDone, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func (pt *PostgresTodo) GetByID(id string) (*Todo, error) {
	var todo Todo
	row := pt.db.QueryRow(`
	SELECT id, description, is_done, created_at, updated_at FROM todos WHERE id = $1`, id)
	err := row.Scan(&todo.ID, &todo.Description, &todo.IsDone, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil{
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrTodoNotFound
		}
		return nil, err
	}
	return &todo, nil
}
