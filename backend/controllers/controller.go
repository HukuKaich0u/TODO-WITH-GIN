package controllers

import (
	"net/http"
	"todo-with-gin/database"
	"todo-with-gin/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "タスクの追加に失敗"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func GetSpecificTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "投稿が見つかりません",
		})
	}
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "タスクの削除に失敗",
		})
	}
}
