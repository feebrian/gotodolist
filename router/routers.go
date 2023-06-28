package router

import (
	"github.com/feebrian/gotodolist/app/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/", controllers.GetTodos)
	r.GET("/todos", controllers.GetTodos)
	r.GET("/todos/:id", controllers.GetTodo)
	r.POST("/todos/create", controllers.CreateTodo)
	r.PATCH("/todos/update/:id", controllers.UpdateTodo)
	r.DELETE("/todos/delete/:id", controllers.DeleteTodo)

	return r
}
