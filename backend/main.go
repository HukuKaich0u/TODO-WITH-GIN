package main

import (
	"todo-with-gin/controllers"
	"todo-with-gin/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.ConnectDB()

	router.POST("/todos", controllers.CreateTodo)
	router.GET("/todos", controllers.GetTodos)

	router.Run("localhost:8080")

}
