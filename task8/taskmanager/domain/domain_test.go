package domain

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)
 type  UserTestSuite struct {		
   suite.Suite
  user User

 }
type  TaskTestSuite  struct{
	suite.Suite
	task Task
}
func (r*UserTestSuite) SetupTest() {
	r.user = User{
		Name: "John Doe",
		Email:"jhon17@gmail.com",
		Password: "password123",
	}
	

}
func (r*UserTestSuite)TearDownTest(){
	fmt.Println("TearDownTest")
}
func (r*TaskTestSuite) TearDownTest() {
	fmt.Println("TearDownTest yo")	
}	
func (r*TaskTestSuite) SetupTest() {
	r.task= Task {
	Title: "Task 1",
	Description: "This is task 1",
	Status: "Pending",
	DueDate: time.Now(),
}

	
}
func (r*UserTestSuite)  Testuser(){
	
	assert.Equal(r.T(), "John Doe", r.user.Name)
	assert.Equal(r.T(), "jhon17@gmail.com", r.user.Email)
	 assert.Equal(r.T(), "password123", r.user.Password)


}
func(r*TaskTestSuite)  Testtask() {

	assert.Equal(r.T(), "Task 1", r.task.Title)
	assert.Equal(r.T(), "This is task 1", r.task.Description)
	assert.Equal(r.T(), "Pending", r.task.Status)
	 assert.Equal(r.T(), time.Now().Format("2006-01-02"), r.task.DueDate.Format("2006-01-02"))


}
func TestDomain(t*testing.T){
	suite.Run(t,new(UserTestSuite))
	suite.Run(t,new(TaskTestSuite))
}    