package domain

import (
	"context"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type  Task  struct{

	ID   primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Status      string `json:"status" bson:"status"`
	DueDate	    time.Time `json:"due_date,omitempty" bson:"due_date,omitempty"`
}
type TaskReipository interface{
	CreateTasks(ctx  context.Context,task Task)(*mongo.InsertOneResult,error)
	DeleteTasks(ctx context.Context,id primitive.ObjectID)  (Task)
    UpdateTask(ctx context.Context, id primitive.ObjectID, updated Task) (*Task, error)
	GetAllTasks(ctx context.Context) ([]Task ,error)
	GetTaskByID(ctx context.Context, id primitive.ObjectID) (*Task, error)


}
type  TaskUsecase  interface{

	CreateTasks(ctx  context.Context,task  Task)(*mongo.InsertOneResult,error)
	DeleteTasks(ctx context.Context, id primitive.ObjectID) (Task)
    UpdateTask(ctx context.Context, id primitive.ObjectID, updated Task) (*Task, error)
	GetAllTasks(ctx context.Context) ([]Task ,error)
	GetTaskByID(ctx context.Context, id primitive.ObjectID) (*Task, error)


}
type User struct{
	Name  string  `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	
}



type  UserUsecase interface{
	Loginfunctionality (ctx context.Context,userinfo User) ( booleanvalue bool, returnedstring string)
	CreateUser(ctx context.Context, user  User) string 


} 
type  UserRepository interface{
	Loginfunctionality (ctx context.Context, userinfo  User) ( booleanvalue bool, returnedstring string)
	CreateUser(ctx context.Context,user  User) string 
}