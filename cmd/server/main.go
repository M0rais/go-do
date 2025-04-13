package main

import (
	"github.com/gin-gonic/gin"
	"github.com/m0rais/go-do/internal/db"
	"github.com/m0rais/go-do/internal/router"
)

func main() {
	db.InitDB()

	server := gin.Default()

	router.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
