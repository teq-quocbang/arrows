// Code generated by mockery v2.36.0. DO NOT EDIT.

package account

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/teq-quocbang/arrows/model"

	uuid "github.com/google/uuid"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// CreateAccount provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) CreateAccount(_a0 context.Context, _a1 *model.Account) (uuid.UUID, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) (uuid.UUID, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) uuid.UUID); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Account) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_CreateAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAccount'
type MockRepository_CreateAccount_Call struct {
	*mock.Call
}

// CreateAccount is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *model.Account
func (_e *MockRepository_Expecter) CreateAccount(_a0 interface{}, _a1 interface{}) *MockRepository_CreateAccount_Call {
	return &MockRepository_CreateAccount_Call{Call: _e.mock.On("CreateAccount", _a0, _a1)}
}

func (_c *MockRepository_CreateAccount_Call) Run(run func(_a0 context.Context, _a1 *model.Account)) *MockRepository_CreateAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Account))
	})
	return _c
}

func (_c *MockRepository_CreateAccount_Call) Return(ID uuid.UUID, err error) *MockRepository_CreateAccount_Call {
	_c.Call.Return(ID, err)
	return _c
}

func (_c *MockRepository_CreateAccount_Call) RunAndReturn(run func(context.Context, *model.Account) (uuid.UUID, error)) *MockRepository_CreateAccount_Call {
	_c.Call.Return(run)
	return _c
}

// GetAccountByConstraint provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) GetAccountByConstraint(_a0 context.Context, _a1 *model.Account) (*model.Account, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) (*model.Account, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) *model.Account); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Account) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetAccountByConstraint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccountByConstraint'
type MockRepository_GetAccountByConstraint_Call struct {
	*mock.Call
}

// GetAccountByConstraint is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *model.Account
func (_e *MockRepository_Expecter) GetAccountByConstraint(_a0 interface{}, _a1 interface{}) *MockRepository_GetAccountByConstraint_Call {
	return &MockRepository_GetAccountByConstraint_Call{Call: _e.mock.On("GetAccountByConstraint", _a0, _a1)}
}

func (_c *MockRepository_GetAccountByConstraint_Call) Run(run func(_a0 context.Context, _a1 *model.Account)) *MockRepository_GetAccountByConstraint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Account))
	})
	return _c
}

func (_c *MockRepository_GetAccountByConstraint_Call) Return(_a0 *model.Account, _a1 error) *MockRepository_GetAccountByConstraint_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetAccountByConstraint_Call) RunAndReturn(run func(context.Context, *model.Account) (*model.Account, error)) *MockRepository_GetAccountByConstraint_Call {
	_c.Call.Return(run)
	return _c
}

// GetAccountByUsername provides a mock function with given fields: ctx, username
func (_m *MockRepository) GetAccountByUsername(ctx context.Context, username string) (*model.Account, error) {
	ret := _m.Called(ctx, username)

	var r0 *model.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Account, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Account); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetAccountByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccountByUsername'
type MockRepository_GetAccountByUsername_Call struct {
	*mock.Call
}

// GetAccountByUsername is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
func (_e *MockRepository_Expecter) GetAccountByUsername(ctx interface{}, username interface{}) *MockRepository_GetAccountByUsername_Call {
	return &MockRepository_GetAccountByUsername_Call{Call: _e.mock.On("GetAccountByUsername", ctx, username)}
}

func (_c *MockRepository_GetAccountByUsername_Call) Run(run func(ctx context.Context, username string)) *MockRepository_GetAccountByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_GetAccountByUsername_Call) Return(_a0 *model.Account, _a1 error) *MockRepository_GetAccountByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetAccountByUsername_Call) RunAndReturn(run func(context.Context, string) (*model.Account, error)) *MockRepository_GetAccountByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// GetList provides a mock function with given fields: _a0
func (_m *MockRepository) GetList(_a0 context.Context) ([]model.Account, error) {
	ret := _m.Called(_a0)

	var r0 []model.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.Account, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.Account); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type MockRepository_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockRepository_Expecter) GetList(_a0 interface{}) *MockRepository_GetList_Call {
	return &MockRepository_GetList_Call{Call: _e.mock.On("GetList", _a0)}
}

func (_c *MockRepository_GetList_Call) Run(run func(_a0 context.Context)) *MockRepository_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRepository_GetList_Call) Return(_a0 []model.Account, _a1 error) *MockRepository_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetList_Call) RunAndReturn(run func(context.Context) ([]model.Account, error)) *MockRepository_GetList_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
