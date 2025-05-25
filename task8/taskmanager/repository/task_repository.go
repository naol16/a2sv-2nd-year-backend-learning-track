package repository

import (
	"context"
	"fmt"
	"log"
	"taskmanager/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	//here I have to import the mongo of client
)
type  taskrepository  struct{
	database  mongo.Database
	collection mongo.Collection
}

func  NewTaskRepo (db mongo.Database,collection mongo.Collection ) domain.TaskReipository{
	return &taskrepository{
		database: db,
        collection: collection,
	}


}
func(r*taskrepository) CreateTasks(ctx  context.Context,task domain.Task)(*mongo.InsertOneResult,error){

return r.collection.InsertOne(ctx,task)
}
func(r*taskrepository) DeleteTasks(ctx context.Context,id primitive.ObjectID) (domain.Task){
	var deleted domain.Task
    filter:=bson.M{"_id":id}
	err:=r.collection.FindOneAndDelete(context.TODO(),filter).Decode(&deleted)
	if err != nil {
        if err == mongo.ErrNoDocuments {
            // No document found for deletion
            return domain.Task{}
        }
        log.Fatal(err)
        return domain.Task{}
    }
	return  deleted







}
func (r*taskrepository)   UpdateTask(ctx context.Context, id primitive.ObjectID, updated domain.Task) (*domain.Task, error){
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
    result, err := r.collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return nil, err
    }
    
    // Check if any document was modified
    if result.ModifiedCount == 0 {
        return nil, fmt.Errorf("no document was updated")
    }
    
    // Fetch and return the updated document
    var updatedTask domain.Task
    err = r.collection.FindOne(context.TODO(), filter).Decode(&updatedTask)
    if err != nil {
        return nil, err
    }
    
    return &updatedTask, nil

    //

}
func(r*taskrepository)GetAllTasks(ctx context.Context) ([]domain.Task ,error){
	findOptions := options.Find()

	
	// Here's an array in which you can store the decoded documents
	var tasks []domain.Task
	
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err :=r.collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		
		// create a value into which the single document can be decoded
		var elem domain.Task
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
func(r*taskrepository)GetTaskByID(ctx context.Context, id primitive.ObjectID) (*domain.Task, error){

 var result*domain.Task
	
	filter:=bson.M{"_id": id}
	err :=r.collection.FindOne(context.TODO(),filter).Decode(&result)
	if err!=nil{
		log.Fatal(err)
	}
	return result,err
}




