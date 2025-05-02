package router

import (
	"taskmanager/controller"
	"taskmanager/middleware"

	"github.com/gin-gonic/gin"
)
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", middleware.AuthMiddleware(),middleware.AuthorizeRole("user"),controller.GetTasks)
	router.GET("/tasks/:id",middleware.AuthMiddleware(),middleware.AuthorizeRole("user"),controller.GetTaskByID)
	router.POST("/tasks", middleware.AuthMiddleware(),middleware.AuthorizeRole("admin"),controller.CreateTask)
	router.PUT("/tasks/:id",middleware.AuthMiddleware(),middleware.AuthorizeRole("admin"),controller.UpdateTask)
	router.DELETE("/tasks/:id", middleware.AuthMiddleware(),middleware.AuthorizeRole("admin"),controller.DeleteTask)
	router.POST("/tasks/user/register",controller.CreateUser)
	router.POST("/tasks/user/login",controller.Loginuser)

	return router
}