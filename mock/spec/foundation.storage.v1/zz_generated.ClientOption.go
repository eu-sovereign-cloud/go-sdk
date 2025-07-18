// Code generated by mockery v2.52.2. DO NOT EDIT.

package mockstorage

import (
	storage "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.storage.v1"
	mock "github.com/stretchr/testify/mock"
)

// MockClientOption is an autogenerated mock type for the ClientOption type
type MockClientOption struct {
	mock.Mock
}

type MockClientOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClientOption) EXPECT() *MockClientOption_Expecter {
	return &MockClientOption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MockClientOption) Execute(_a0 *storage.Client) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*storage.Client) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClientOption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockClientOption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *storage.Client
func (_e *MockClientOption_Expecter) Execute(_a0 interface{}) *MockClientOption_Execute_Call {
	return &MockClientOption_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MockClientOption_Execute_Call) Run(run func(_a0 *storage.Client)) *MockClientOption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*storage.Client))
	})
	return _c
}

func (_c *MockClientOption_Execute_Call) Return(_a0 error) *MockClientOption_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClientOption_Execute_Call) RunAndReturn(run func(*storage.Client) error) *MockClientOption_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClientOption creates a new instance of MockClientOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClientOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClientOption {
	mock := &MockClientOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
