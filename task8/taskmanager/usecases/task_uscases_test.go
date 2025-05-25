package usecases_test

import (
	"context"
	"fmt"
	"taskmanager/domain"
	"taskmanager/mocks"
	"taskmanager/usecases"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

type  TaskusecaseTestSuite struct {
	suite.Suite
	repository  *mocks.TaskReipository
	TaskUsecase domain.TaskUsecase

}

func (r *TaskusecaseTestSuite) SetupTest() {
	r.repository = new(mocks.TaskReipository)
	r.TaskUsecase = usecases.NewTaskUsecase(r.repository, 10*time.Second)
}

func (r *TaskusecaseTestSuite) TestDeleteTasks() {
	taskID := "60c72b2f9b1d8c001c8e4f5a" // Example task ID
	id, _ :=primitive.ObjectIDFromHex(taskID)
	deletedTask := domain.Task{
		ID: id,
		Title: "Sample Task",
		Description: "This is a sample task",
		Status: "Pending",
	}

	r.repository.On("DeleteTasks", mock.Anything, id).Return(deletedTask).Once()
	result := r.TaskUsecase.DeleteTasks(context.Background(), id)
	assert.Equal(r.T(), deletedTask, result, "Expected deleted task to match")	

}
func (r *TaskusecaseTestSuite) TestUpdateTask() {
	taskID := "60c72b2f9b1d8c001c8e4f5a" // Example task ID		

	id, _ := primitive.ObjectIDFromHex(taskID)		
	updatedTask := domain.Task{
		ID: id,	
		Title: "Updated Task",	
		Description: "This is an updated task",
		Status: "Completed",
	}
	r.repository.On("UpdateTask", mock.Anything, id, updatedTask).Return(&updatedTask, nil).Once()	
	result, err := r.TaskUsecase.UpdateTask(context.Background(), id, updatedTask)
	assert.Nil(r.T(), err, "Expected no error when updating task")
	assert.Equal(r.T(), &updatedTask, result, "Expected updated task to match")
}
func (r *TaskusecaseTestSuite) TestGetAllTasks() {
	tasks := []domain.Task{{
		ID: primitive.NewObjectID(),
		Title: "Task 1",
		Description: "Description for Task 1",
		Status: "Pending",
	}, {
		ID: primitive.NewObjectID(),
		Title: "Task 2",
		Description: "Description for Task 2",
		Status: "Completed",

	}}
	r.repository.On("GetAllTasks", mock.Anything).Return(tasks, nil).Once()
	result, err := r.TaskUsecase.GetAllTasks(context.Background())
	assert.Nil(r.T(), err, "Expected no error when getting all tasks")
	assert.Equal(r.T(), tasks, result, "Expected all tasks to match")
		}
func (r *TaskusecaseTestSuite) TestGetTaskByID() {
	taskID := "60c72b2f9b1d8c001c8e4f5a" // Example task ID	
	id, _ := primitive.ObjectIDFromHex(taskID)
	task := &domain.Task{
		ID: id,
		Title: "Sample Task",
		Description: "This is a sample task",
		Status: "Pending",	
	}
	r.repository.On("GetTaskByID", mock.Anything, id).Return(task, nil).Once()
	result, err := r.TaskUsecase.GetTaskByID(context.Background(), id)
	assert.Nil(r.T(), err, "Expected no error when getting task by ID")
	assert.Equal(r.T(), task, result, "Expected task to match")
}
func (r*TaskusecaseTestSuite) TearDownTest(){
	r.repository.AssertExpectations(r.T())
	fmt.Println("teardown test called")
}

func TestTaskusecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskusecaseTestSuite))
}	