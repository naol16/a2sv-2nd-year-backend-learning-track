package repository

import ( "taskmanager/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"taskmanager/infrastrcure"

	"context"
	"fmt"
	"log"

)

type userepository struct {
	database   mongo.Database
	collection mongo.Collection
}

func 	NewUserRepo(db mongo.Database, collection mongo.Collection) domain.UserRepository{
	return &userepository{
		database:db,
		collection:collection,
	}
}

func(r*userepository)Loginfunctionality (ctx context.Context,userinfo  domain.User ) ( booleanvalue bool, returnedstring string){
var  user domain.User
type newuser  struct{
	user  domain.User
}
	filter := bson.M{"email":userinfo.Email}
	// and then cross check 
	//do the thing
	err :=r.collection.FindOne(context.TODO(),filter).Decode(&user)
	if err==mongo.ErrNoDocuments{
		return false, "user does not exist"
	}
	usersHashedpassword := user.Password
   // here we will use the function
    value:= infrastrcure.LoginChekcer(userinfo, usersHashedpassword)
	if value!=" "{
		return false, value
	}


value ,error := infrastrcure.Generator(user)
if error!=nil{
	log.Fatal(error)
}
return true,value

}
func(r*userepository) CreateUser(ctx context.Context, user  domain.User) string {

	 var existingUser domain.User
    err := r.collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
    if err == nil {
        return "email already exists"
    }
    // Only proceed if the error is "no documents" (meaning email doesn't exist)
    if err != nil && err != mongo.ErrNoDocuments {
        log.Fatal(err)
        return "error checking email availability"
    }

    password := user.Password
    hashedpassword := infrastrcure.Hasher(password)
	
    user.Password = string(hashedpassword)
    result, err := r.collection.InsertOne(context.TODO(), user)
    if err != nil {
        log.Fatal(err)
        return "there is error while creating a user"
    }
    fmt.Println("inserted id", result.InsertedID)
    return "user created"

	}

