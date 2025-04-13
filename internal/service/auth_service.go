package service

import (
	"errors"
	"github.com/m0rais/go-do/internal/models"
	"github.com/m0rais/go-do/internal/repositories"
	"github.com/m0rais/go-do/internal/utils"
)

func ValidateCredentials(input *models.User) error {
	user, err := repositories.GetUserByEmail(input.Email)
	if err != nil {
		return err
	}

	if !utils.CompareHashAndPassword(user.Password, input.Password) {
		return errors.New("Invalid credentials")
	}

	return nil
}
