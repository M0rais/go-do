package repositories

import (
	"github.com/m0rais/go-do/internal/db"
	"github.com/m0rais/go-do/internal/models"
	"github.com/m0rais/go-do/internal/utils"
)

func SaveUser(user *models.User) error {

	password, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = password

	query := `
        INSERT INTO users (email, password)
        VALUES (:email, :password)
    `

	result, err := db.DB.NamedExec(query, user)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	user.Id = id

	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
