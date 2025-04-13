package models

import (
	"time"
)

type Todo struct {
	Id          int64     `db:"id"`
	Name        string    `db:"name" binding:"required"`
	Description string    `db:"description" binding:"required"`
	StartDate   time.Time `db:"startDate" binding:"required"`
	EndDate     time.Time `db:"endDate" binding:"required"`
	Completed   bool      `db:"completed"`
	UserId      int64     `db:"userId"`
}

func NewTodo(name string, description string, startDate time.Time, endDate time.Time) *Todo {
	return &Todo{
		Name:        name,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
		Completed:   false,
	}
}
