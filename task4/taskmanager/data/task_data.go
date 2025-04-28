package data

import (
	"context"
	"fmt"
	"log"
	 "os"

	"taskmanager/model"
	"time"

    "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

var tasks = []model.Task{
	{ Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

var collection*mongo.Collection

func Initalizationconnection() mongo.Client{
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
    collection = client.Database("gobackend").Collection("todotask")
	
	return *client
	}



func GetAllTasks() ([]model.Task ,error){

	findOptions := options.Find()

	
	// Here's an array in which you can store the decoded documents
	var tasks []model.Task
	
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		
		// create a value into which the single document can be decoded
		var elem model.Task
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
	
	tasks= append(tasks, elem)
	}
	
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	
	// Close the cursor once finished
	defer cur.Close(context.TODO())
	return tasks,nil
	

}

func GetTaskByID(id primitive.ObjectID) (*model.Task, error) {
   var result*model.Task
	
	filter:=bson.M{"_id": id}
	err :=collection.FindOne(context.TODO(),filter).Decode(&result)
	if err!=nil{
		log.Fatal(err)
	}
	return result,err
}

func CreateTask(task model.Task) (*mongo.InsertOneResult, error) {

	return collection.InsertOne(context.TODO(),task)
	
}

func UpdateTask(id primitive.ObjectID, updated model.Task) (*model.Task, error) {
	filter := bson.M{"_id": id}
    
    // Create an update document with $set operator
    update := bson.M{
        "$set": bson.M{},
    }
    
    // Add fields to update dynamically
    if updated.Title != "" {
        update["$set"].(bson.M)["title"] = updated.Title
    }
    if updated.Description != "" {
        update["$set"].(bson.M)["description"] = updated.Description
    }
    if !updated.DueDate.IsZero() { // Check if DueDate is set
        update["$set"].(bson.M)["dueDate"] = updated.DueDate
    }
    if updated.Status != " " {
        update["$set"].(bson.M)["completed"] = updated.Status
    }
    // Add other fields as needed
    
    // Perform the update
    result, err := collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return nil, err
    }
    
    // Check if any document was modified
    if result.ModifiedCount == 0 {
        return nil, fmt.Errorf("no document was updated")
    }
    
    // Fetch and return the updated document
    var updatedTask model.Task
    err = collection.FindOne(context.TODO(), filter).Decode(&updatedTask)
    if err != nil {
        return nil, err
    }
    
    return &updatedTask, nil

    // Build the update data dynamically using the request body
    
	
}

func DeleteTask(id primitive.ObjectID) model.Task {
	var deleted model.Task
    filter:=bson.M{"_id":id}
	err:=collection.FindOneAndDelete(context.TODO(),filter).Decode(&deleted)
	if err != nil {
        if err == mongo.ErrNoDocuments {
            // No document found for deletion
            return model.Task{}
        }
        log.Fatal(err)
        return model.Task{}
    }

    return deleted
	
   


}

