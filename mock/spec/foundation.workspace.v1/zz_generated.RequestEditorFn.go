// Code generated by mockery v2.52.2. DO NOT EDIT.

package mockworkspace

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockRequestEditorFn is an autogenerated mock type for the RequestEditorFn type
type MockRequestEditorFn struct {
	mock.Mock
}

type MockRequestEditorFn_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRequestEditorFn) EXPECT() *MockRequestEditorFn_Expecter {
	return &MockRequestEditorFn_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, req
func (_m *MockRequestEditorFn) Execute(ctx context.Context, req *http.Request) error {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRequestEditorFn_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockRequestEditorFn_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - req *http.Request
func (_e *MockRequestEditorFn_Expecter) Execute(ctx interface{}, req interface{}) *MockRequestEditorFn_Execute_Call {
	return &MockRequestEditorFn_Execute_Call{Call: _e.mock.On("Execute", ctx, req)}
}

func (_c *MockRequestEditorFn_Execute_Call) Run(run func(ctx context.Context, req *http.Request)) *MockRequestEditorFn_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockRequestEditorFn_Execute_Call) Return(_a0 error) *MockRequestEditorFn_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequestEditorFn_Execute_Call) RunAndReturn(run func(context.Context, *http.Request) error) *MockRequestEditorFn_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRequestEditorFn creates a new instance of MockRequestEditorFn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRequestEditorFn(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRequestEditorFn {
	mock := &MockRequestEditorFn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
