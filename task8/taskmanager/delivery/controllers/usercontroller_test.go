package controllers_test

import (
	"taskmanager/delivery/controllers"
	"taskmanager/domain/mocks"
	"taskmanager/utils"

	// "github.com/stretchr/testify/mock"
	// "github.com/stretchr/testify/assert"
	"bytes"
	"net/http"
	"taskmanager/domain"

	"encoding/json"
	"testing"

	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)


type UserControllerTestSuite struct{
	controller   controllers.UserController
	usercase    *mocks.UserUsecase
    suite.Suite
	testingServer   *httptest.Server
    

}
func (suite *UserControllerTestSuite) SetupTest() {
	suite.usercase= new(mocks.UserUsecase)

	suite.controller = controllers.UserController{
		Usecase: suite.usercase,	
	}
	router :=gin.Default();
	router.POST("/users", utils.HTTP(suite.controller.CreateUser))
	router.POST("/login", utils.HTTP(suite.controller.Loginuser))
	suite.testingServer = httptest.NewServer(router)


}
func (suite *UserControllerTestSuite) TestCreateUser() {
	user := domain.User{
		Email:  "checkemail",
		Name:"name",
		Password: "password",
		Role: "user",
	}
	suite.usercase.On("CreateUser",&user).Return("user created").Once()
   requestbody, err := json.Marshal(user)
   if err != nil {
		suite.T().Fatalf("Error marshalling user: %v", err)
   }
	response, _ := http.Post(suite.testingServer.URL+"/users", "application/json", bytes.NewBuffer(requestbody))
	// resp, err := suite.testingServer.Client().Do(re
	suite.Equal(http.StatusCreated, response.StatusCode)
	assert.Equal(suite.T(), "user created", response.Status)
	
	suite.usercase.AssertExpectations(suite.T())
}


func (suite *UserControllerTestSuite) TearDownTest() {
	suite.usercase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestLoginUser() {
	user := domain.User{
		Email:  "checkemail",
		Password: "password",
	}
	suite.usercase.On("LoginUser", &user).Return(true, "token").Once()
	requestbody, err := json.Marshal(user)
	if err != nil {
		suite.T().Fatalf("Error marshalling user: %v", err)
	}
	response, _ := http.Post(suite.testingServer.URL+"/login", "application/json", bytes.NewBuffer(requestbody))
	suite.Equal(http.StatusOK, response.StatusCode)
	assert.Equal(suite.T(), "token", response.Status)
	
	suite.usercase.AssertExpectations(suite.T())
}	
func TestUsercontroller(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}