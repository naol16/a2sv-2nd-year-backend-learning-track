package main
import(
"taskmanager/router"
)
func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
