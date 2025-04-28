package main

import (
	"context"
	"log"
	"taskmanager/data"
	"taskmanager/router"
)
func main() {
	// fmt.Println("MongoDB URI:", os.Getenv(MONGO_URI))



// Close the connection when the main function exit

client :=data.Initalizationconnection()
r := router.SetupRouter()
r.Run(":8081")
if err := client.Disconnect(context.TODO()); err != nil {
			log.Printf("Warning: failed to disconnect from MongoDB: %v", err)
		} else {
			log.Println("MongoDB connection closed")
		}
}
