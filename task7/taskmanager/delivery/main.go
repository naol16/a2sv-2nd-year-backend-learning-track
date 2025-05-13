package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"taskmanager/delivery/routers"
	"time"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func main() {
	// fmt.Println("MongoDB URI:", os.Getenv(MONGO_URI))

//here it is the context
ctx := context.TODO()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
    myurl:=os.Getenv("MONGO_URI")
	clientOptions:= options.Client().ApplyURI(myurl)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Mongo connection error:", err)
	}
	
	
	if err := client.Ping(ctx, nil)
	err != nil {
		log.Fatal("Ping error:", err)
	}
	
	fmt.Println("Connected to MongoDB!")
	  database := client.Database("gobackend")
	 taskcollection := client.Database("gobackend").Collection("todotask")
	usercollection  := client.Database("gobackend").Collection("user")


	timeout :=   10*time.Second

   r := routers.SetupRouter(*database,*taskcollection,*usercollection,timeout)
	r.Use(gin.Logger())
   r.Run(":8081")
if err := client.Disconnect(context.TODO()); err != nil {
			log.Printf("Warning: failed to disconnect from MongoDB: %v", err)
		} else {
			log.Println("MongoDB connection closed")
		}
}
