package controller

import (
	"log"
	"net/http"
	"taskmanager/data"
	"taskmanager/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

)
type Task struct{
	task model.Task 
}


func  CreateUser(c*gin.Context){
	var user  model.User
	err :=c.BindJSON(&user)
	if err!=nil{
		log.Fatal(err)

	}
	result := data.CreateUser(user)
  c.JSON(http.StatusCreated,result)



}

func  Loginuser(c*gin.Context){
  var  userinfo model.User
  err:= c.BindJSON(&userinfo)
  if err!=nil{
	log.Fatal(err)
  }
  booleanvalue,messagevalue := data.Loginfunctionality(userinfo)
  if booleanvalue==false{
	c.JSON(http.StatusBadRequest,messagevalue)

  }else{
	
	c.JSON(http.StatusAccepted,messagevalue)
  }
}
func GetTasks(c *gin.Context) {
	tasks,err:=data.GetAllTasks()
	if err!=nil{
		log.Fatal(err)
	}

	c.JSON(http.StatusOK,tasks)
}

 func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err!=nil{
		log.Fatal(err)
	}

	 task, err := data.GetTaskByID(objectID)
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
	created, err := data.CreateTask(task)
if err != nil {
    // Handle error (log, return HTTP error, etc.)
    log.Printf("Failed to create task: %v", err)
    return // or c.JSON() if in a Gin handler
}
	c.JSON(http.StatusCreated, created)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err!=nil{
		log.Fatal(err)
	}



	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := data.UpdateTask(objectID,task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func DeleteTask(c *gin.Context) {
	id:=c.Param("id")
	objectid, err := primitive.ObjectIDFromHex(id)
	if  err!=nil {
		log.Fatal(err)
	}


	 deleted := data.DeleteTask(objectid); 
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		
	}

	c.JSON(http.StatusOK, gin.H{"deleted":deleted})
}