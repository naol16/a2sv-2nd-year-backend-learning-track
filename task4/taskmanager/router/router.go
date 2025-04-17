package router
import(
	"github.com/gin-gonic/gin"
	"taskmanager/controller"
)
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controller.GetTasks)
	router.GET("/tasks/:id", controller.GetTaskByID)
	router.POST("/tasks", controller.CreateTask)
	router.PUT("/tasks/:id", controller.UpdateTask)
	router.DELETE("/tasks/:id", controller.DeleteTask)

	return router
}