package controller

import (
	"net/http"

	"taskmanager/data"
	"taskmanager/model"

	"github.com/gin-gonic/gin"
)
type Task struct{
	task model.Task 
}

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, data.GetAllTasks())
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created := data.CreateTask(task)
	c.JSON(http.StatusCreated, created)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")


	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := data.UpdateTask(id,task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func DeleteTask(c *gin.Context) {
	id:=c.Param("id")

	if err := data.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}