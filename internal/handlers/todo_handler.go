package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/m0rais/go-do/internal/models"
	"github.com/m0rais/go-do/internal/repositories"
	"net/http"
	"strconv"
)

func GetTodos(context *gin.Context) {
	userId := context.MustGet("userID").(int64)

	todos, err := repositories.GetAllTodos(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, todos)
}

func GetTodo(context *gin.Context) {
	userId := context.MustGet("userID").(int64)
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	todo, err := repositories.GetTodoById(id, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, todo)
}

func CreateTodo(context *gin.Context) {
	userId := context.MustGet("userID").(int64)

	var todo models.Todo
	err := context.ShouldBindJSON(&todo)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	todo.UserId = userId

	err = repositories.SaveTodo(&todo)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusCreated, todo)
}

func CompleteTodo(context *gin.Context) {
	userId := context.MustGet("userID").(int64)

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	todo, err := repositories.GetTodoById(id, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = repositories.CompleteTodo(todo)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	todo.Completed = true

	context.JSON(http.StatusOK, todo)
}
