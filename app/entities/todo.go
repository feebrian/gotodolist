package entities

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primary_key"`
	UserID      uint   `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Completed   bool   `json:"completed"`
}

type CreateTodo struct {
	gorm.Model

	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	DueDate     string `json:"dueDate" binding:"required"`
}

type UpdateTodo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Completed   bool   `json:"completed"`
}
