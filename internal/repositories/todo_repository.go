package repositories

import (
	"github.com/m0rais/go-do/internal/db"
	"github.com/m0rais/go-do/internal/models"
)

func SaveTodo(todo *models.Todo) error {
	query := `
        INSERT INTO todos (name, description, startDate, endDate, completed, userId)
        VALUES (:name, :description, :startDate, :endDate, :completed, :userId)
    `

	result, err := db.DB.NamedExec(query, todo)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err == nil {
		todo.Id = id
	}
	return err
}

func CompleteTodo(todo *models.Todo) error {
	query := `UPDATE todos SET completed = 1 WHERE id = :id`
	_, err := db.DB.NamedExec(query, todo)
	return err
}

func GetAllTodos(userId int64) ([]models.Todo, error) {
	var todos []models.Todo
	err := db.DB.Select(&todos, "SELECT * FROM todos WHERE userId = ?", userId)
	return todos, err
}

func GetTodoById(id int64, userId int64) (*models.Todo, error) {
	var todo models.Todo
	err := db.DB.Get(&todo, "SELECT * from todos WHERE id = ? AND userId = ?", id, userId)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}
