package main
import(
 "library_management/controllers"
 "library_management/services"
)
func main(){
	 // Initialize service
	 library := services.NewLibrary()
    
	 // Create controller
	 controller := controllers.NewController(library)  
	 controller.Run()
}