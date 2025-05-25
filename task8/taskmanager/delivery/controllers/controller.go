package controllers


import (
	"log"
	"net/http"

	"taskmanager/domain"
	

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

)
type  UserController struct {
	Usecase   domain.UserUsecase
}
type TaskController  struct{
	Taskusecases  domain.TaskUsecase

}

func (r*UserController) CreateUser(c*gin.Context){
	var user domain.User
	err :=c.BindJSON(&user)
	if err!=nil{
		log.Fatal(err)

	}
 	result := r.Usecase.CreateUser(c,user)
  c.JSON(http.StatusCreated,result)



}

func (r*UserController) Loginuser(c*gin.Context){
  var  userinfo domain.User
  err:= c.BindJSON(&userinfo)
  if err!=nil{
	log.Fatal(err)
  }
  booleanvalue,messagevalue :=r.Usecase.Loginfunctionality(c,userinfo)
  if booleanvalue==false{
	c.JSON(http.StatusBadRequest,messagevalue)

  }else{
	
	c.JSON(http.StatusAccepted,messagevalue)
  }
}
func(r*TaskController) GetTasks(c *gin.Context) {
	// here the  context  is the error 
	tasks,err:=r.Taskusecases.GetAllTasks(c)
	if err!=nil{
		log.Fatal(err)
	}

	c.JSON(http.StatusOK,tasks)
}

 func(r*TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err!=nil{
		log.Fatal(err)
	}

	 task, err := r.Taskusecases.GetTaskByID(c,objectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func(r*TaskController)CreateTask(c *gin.Context) {
	var task  domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	
	created, err := r.Taskusecases.CreateTasks(c,task)
if err != nil {
    // Handle error (log, return HTTP error, etc.)
    log.Printf("Failed to create task: %v", err)
    return // or c.JSON() if in a Gin handler
}
	c.JSON(http.StatusCreated, created)
}

func (r*TaskController)UpdateTask(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err!=nil{
		log.Fatal(err)
	}



	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err :=r.Taskusecases.UpdateTask(c,objectID,task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}//

func(r*TaskController) DeleteTask(c *gin.Context) {
	id:=c.Param("id")
	objectid, err := primitive.ObjectIDFromHex(id)
	if  err!=nil {
		log.Fatal(err)
	}


	 deleted := r.Taskusecases.DeleteTasks(c,objectid); 
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		
	}

	c.JSON(http.StatusOK, gin.H{"deleted":deleted})
}



