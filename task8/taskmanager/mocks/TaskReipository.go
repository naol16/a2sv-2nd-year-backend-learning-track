// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "taskmanager/domain"

	mock "github.com/stretchr/testify/mock"

	mongo "go.mongodb.org/mongo-driver/mongo"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskReipository is an autogenerated mock type for the TaskReipository type
type TaskReipository struct {
	mock.Mock
}

// CreateTasks provides a mock function with given fields: ctx, task
func (_m *TaskReipository) CreateTasks(ctx context.Context, task domain.Task) (*mongo.InsertOneResult, error) {
	ret := _m.Called(ctx, task)

	if len(ret) == 0 {
		panic("no return value specified for CreateTasks")
	}

	var r0 *mongo.InsertOneResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Task) (*mongo.InsertOneResult, error)); ok {
		return rf(ctx, task)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Task) *mongo.InsertOneResult); ok {
		r0 = rf(ctx, task)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.InsertOneResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Task) error); ok {
		r1 = rf(ctx, task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTasks provides a mock function with given fields: ctx, id
func (_m *TaskReipository) DeleteTasks(ctx context.Context, id primitive.ObjectID) domain.Task {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTasks")
	}

	var r0 domain.Task
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) domain.Task); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	return r0
}

// GetAllTasks provides a mock function with given fields: ctx
func (_m *TaskReipository) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllTasks")
	}

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Task, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Task); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTaskByID provides a mock function with given fields: ctx, id
func (_m *TaskReipository) GetTaskByID(ctx context.Context, id primitive.ObjectID) (*domain.Task, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskByID")
	}

	var r0 *domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) (*domain.Task, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *domain.Task); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: ctx, id, updated
func (_m *TaskReipository) UpdateTask(ctx context.Context, id primitive.ObjectID, updated domain.Task) (*domain.Task, error) {
	ret := _m.Called(ctx, id, updated)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTask")
	}

	var r0 *domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, domain.Task) (*domain.Task, error)); ok {
		return rf(ctx, id, updated)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, domain.Task) *domain.Task); ok {
		r0 = rf(ctx, id, updated)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID, domain.Task) error); ok {
		r1 = rf(ctx, id, updated)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTaskReipository creates a new instance of TaskReipository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskReipository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskReipository {
	mock := &TaskReipository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
