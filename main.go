package main

import (
	"net/http"
	"todo-with-gin/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.ConnectDB()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	router.Run("localhost:8080")

}
