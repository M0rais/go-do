package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/m0rais/go-do/internal/models"
	"github.com/m0rais/go-do/internal/repositories"
	"github.com/m0rais/go-do/internal/service"
	"github.com/m0rais/go-do/internal/utils"
	"net/http"
)

func Register(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = repositories.SaveUser(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusCreated, user)
}

func Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = service.ValidateCredentials(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}
