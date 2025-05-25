package repository

import (
	"fmt"
	"log"
	"os"
	"taskmanager/domain"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"context"



	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type  userrepositoryTestSuite struct {	
	suite.Suite
	userRepository userepository
    usercollection *mongo.Collection
}	

func  DbInitialization() (*mongo.Database,*mongo.Collection,*mongo.Collection) {
	//here it is the context
err := godotenv.Load()
if err != nil {
	log.Fatal("Error loading .env file")
}
ctx := context.TODO()
	
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
	return database,taskcollection, usercollection

}
func (r*userrepositoryTestSuite) SetupTest() {
	//here it is the context
	database,taskcollection, usercollection :=DbInitialization()
	r.usercollection = usercollection

	if taskcollection==nil{
		log.Fatal("Mongo connection error:")
	}
	r.userRepository = userepository{	
		database:   *database,
		collection: *usercollection,
	}

		
}
func (r*userrepositoryTestSuite) TearDownTest() {
	fmt.Println("TearDownTest")
}
func (r*userrepositoryTestSuite) CreateChecker() {
	user := domain.User{
		Name:     "John Doe",
		Email:    "dsfs@gmail.com",
		Password: "password123",
		Role:     "user",
	}
	ctx := context.TODO()
	value :=r.userRepository.CreateUser(ctx, user)
	assert.Equal(r.T(), "user created successfully", value)
  
}
func (r*userrepositoryTestSuite) LoginChecker() {
	user := domain.User{
		Name:     "John Doe",
		Email:    "dsfs@gmail.com",
		Password: "password123",
}
	ctx := context.TODO()
	value  ,value2:= r.userRepository.Loginfunctionality(ctx, user)
	assert.Equal(r.T(), true, value)
	assert.NotNil(r.T(), value2)

}
func  TestUserRepository(t*testing.T) {
	suite.Run(t, new(userrepositoryTestSuite))
}