package data

import (

	"context"
	"fmt"
	"log"
	"os"

	"taskmanager/model"
	
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	 "github.com/dgrijalva/jwt-go"

)


var usercollection*mongo.Collection
var  taskcollection*mongo.Collection
var client*mongo.Client

func Initalizationconnection() mongo.Client{
	ctx := context.TODO()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
    myurl:=os.Getenv("MONGO_URI")
	clientOptions:= options.Client().ApplyURI(myurl)
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Mongo connection error:", err)
	}
	
	
	if err := client.Ping(ctx, nil)
	err != nil {
		log.Fatal("Ping error:", err)
	}
	
	fmt.Println("Connected to MongoDB!")
    taskcollection= client.Database("gobackend").Collection("todotask")
	usercollection  = client.Database("gobackend").Collection("user")
	
	return *client
	}



func GetAllTasks() ([]model.Task ,error){

	findOptions := options.Find()

	
	// Here's an array in which you can store the decoded documents
	var tasks []model.Task
	
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := taskcollection.Find(context.TODO(), bson.D{{}}, findOptions)
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
func  Loginfunctionality (userinfo  model.User) ( booleanvalue bool, returnedstring string){
	//filter 
	var  user model.User
	filter := bson.M{"email":userinfo.Email}
	// and then cross check 
	//do the thing
	err :=usercollection.FindOne(context.TODO(),filter).Decode(&user)
	if err==mongo.ErrNoDocuments{
		return false, "user does not exist"
	}
	usersHashedpassword := user.Password
	bcrypterr :=bcrypt.CompareHashAndPassword([]byte(usersHashedpassword),[]byte(userinfo.Password))
	if bcrypterr!=nil{
		return false,"either your password or username does not much"
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mysecretkey := os.Getenv("JWT_SECRET")
	fmt.Println(mysecretkey)
	var jwtsecret= []byte (mysecretkey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   user.Email,
		"name": user.Name,
		"role": user.Role,
	  })
	  
	  jwtToken, err := token.SignedString(jwtsecret)

	

   return true ,jwtToken


}

func CreateUser(user model.User) string {
    // First check if email already exists
    var existingUser model.User
    err := usercollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
    if err == nil {
        return "email already exists"
    }
    // Only proceed if the error is "no documents" (meaning email doesn't exist)
    if err != nil && err != mongo.ErrNoDocuments {
        log.Fatal(err)
        return "error checking email availability"
    }

    fmt.Println(usercollection)
    password := user.Password
    hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    if err != nil {
        log.Fatal(err)
    }
    user.Password = string(hashedpassword)
    result, err := usercollection.InsertOne(context.TODO(), user)
    if err != nil {
        log.Fatal(err)
        return "there is error while creating a user"
    }
    fmt.Println("inserted id", result.InsertedID)
    return "user created"
}
func GetTaskByID(id primitive.ObjectID) (*model.Task, error) {
   var result*model.Task
	
	filter:=bson.M{"_id": id}
	err :=taskcollection.FindOne(context.TODO(),filter).Decode(&result)
	if err!=nil{
		log.Fatal(err)
	}
	return result,err
}

func CreateTask(task model.Task) (*mongo.InsertOneResult, error) {

	return taskcollection.InsertOne(context.TODO(),task)
	
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
    result, err := taskcollection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return nil, err
    }
    
    // Check if any document was modified
    if result.ModifiedCount == 0 {
        return nil, fmt.Errorf("no document was updated")
    }
    
    // Fetch and return the updated document
    var updatedTask model.Task
    err = taskcollection.FindOne(context.TODO(), filter).Decode(&updatedTask)
    if err != nil {
        return nil, err
    }
    
    return &updatedTask, nil

    // Build the update data dynamically using the request body
    
	
}

func DeleteTask(id primitive.ObjectID) model.Task {
	var deleted model.Task
    filter:=bson.M{"_id":id}
	err:=taskcollection.FindOneAndDelete(context.TODO(),filter).Decode(&deleted)
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

