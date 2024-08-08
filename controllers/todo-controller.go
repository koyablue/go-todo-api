package controllers

import (
	"encoding/json"
	"go-todo-api/models"
	"net/http"
)

type TodoController struct {
	Model *models.TodoModel
}

func NewTodoController(m *models.TodoModel) *TodoController {
	return &TodoController{Model: m}
}

func (todoController *TodoController) GetTodos(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	todos, err := todoController.Model.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}
