package routers

import (
	"taskmanager/delivery/controllers"
	"taskmanager/infrastrcure"

	"taskmanager/repository"
	"taskmanager/usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)
func SetupRouter(db mongo.Database,taskcollection mongo.Collection,usercollection mongo.Collection, coltimeout time.Duration) *gin.Engine {
	tr:=  repository.NewTaskRepo(db,taskcollection)
	ur:= repository.NewUserRepo(db,usercollection)
	uc := &controllers.UserController{
		Usecase: usecases.NewUserUsecase(ur,coltimeout),
	}
	tc:= &controllers.TaskController{
		Taskusecases : usecases.NewTaskUsecase(tr,coltimeout),
	}
	router:= gin.Default()
	// db,collection

    router.POST("/users",uc.CreateUser)
	router.POST("/login",uc.Loginuser)
	authenticateduserrouter :=router.Group(" ")
	authenticateduserrouter.Use(infrastrcure.AuthMiddleware())
	authenticateduserrouter.Use(infrastrcure.AuthorizeRole("user"))
	authenticateduserrouter.GET("/task",tc.GetTasks)
	authenticateduserrouter.GET("tasks/:id",tc.GetTaskByID)
	authenticatedadminrouter :=router.Group(" ")
	authenticatedadminrouter.Use(infrastrcure.AuthMiddleware())
	authenticatedadminrouter.Use(infrastrcure.AuthorizeRole("admin"))
	authenticatedadminrouter.POST("/tasks",tc.CreateTask)
	authenticatedadminrouter.PUT("/tasks/:id",tc.UpdateTask)
	authenticatedadminrouter.DELETE("/tasks/:id",tc.DeleteTask)
	return router
}







   