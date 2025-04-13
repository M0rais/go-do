package models

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func NewUser(name string, email string, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}
