package models

import "database/sql"

type Todo struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
}

type TodoModel struct {
	DB *sql.DB
}

func NewTodoModel(db *sql.DB) *TodoModel {
	return &TodoModel{DB: db}
}

// Returns all todos
func (todoModel *TodoModel) All() ([]Todo, error) {
	rows, err := todoModel.DB.Query("SELECT id, task FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.Task); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
