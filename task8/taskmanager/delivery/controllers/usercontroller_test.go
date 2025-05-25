package controllers_test

import (
	"taskmanager/delivery/controllers"
	"taskmanager/domain/mocks"
	"taskmanager/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"bytes"
	"taskmanager/domain"
	"net/http"
	"encoding/json"

	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
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
   requestbody,err := json.Marshal(user)
	req, _ := http.Post(suite.testingServer.URL+"/users", "application/json", bytes.NewBuffer(requestbody))
	// resp, err := suite.testingServer.Client().Do(req)
	

	suite.NoError(err)
	suite.Equal(201, resp.StatusCode)
	suite.usercase.AssertExpectations(suite.T())
}


func (suite *UserControllerTestSuite) TearDownTest() {
	suite.usercase.AssertExpectations(suite.T())
}

