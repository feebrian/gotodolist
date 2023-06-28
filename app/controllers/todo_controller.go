package controllers

import (
	"net/http"

	"github.com/feebrian/gotodolist/app/entities"
	"github.com/feebrian/gotodolist/config"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []entities.Todo
	config.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func GetTodo(c *gin.Context) {
	var todo entities.Todo
	var err error

	if err = config.DB.Where("id == ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func CreateTodo(c *gin.Context) {
	var input entities.CreateTodo
	var err error

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := entities.Todo{
		Title:       input.Title,
		Description: input.Description,
		DueDate:     input.DueDate,
	}

	config.DB.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{"message": "todo successfully created"})
}

func UpdateTodo(c *gin.Context) {
	var err error

	// check if todo exist in database
	var todo entities.Todo
	if err = config.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check input to update todo
	var input entities.UpdateTodo
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&todo).First(&todo).Updates(&input)
	c.JSON(http.StatusOK, gin.H{"message": "todo updated sucessfully"})
}

func DeleteTodo(c *gin.Context) {
	var err error

	// check if todo exist in database
	var todo entities.Todo
	if err = config.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if exist then delete the todo
	if err = config.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "sucessfully delete todo"})
}
