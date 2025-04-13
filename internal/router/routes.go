package router

import (
	"github.com/gin-gonic/gin"
	"github.com/m0rais/go-do/internal/handlers"
	"github.com/m0rais/go-do/internal/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	authenticatedRoutes := server.Group("/")
	authenticatedRoutes.Use(middlewares.Authenticate)
	authenticatedRoutes.GET("/todos", handlers.GetTodos)
	authenticatedRoutes.GET("/todos/:id", handlers.GetTodo)
	authenticatedRoutes.POST("/todos", handlers.CreateTodo)
	authenticatedRoutes.PATCH("/todos/:id", handlers.CompleteTodo)

	server.POST("/login", handlers.Login)
	server.POST("/register", handlers.Register)
}
