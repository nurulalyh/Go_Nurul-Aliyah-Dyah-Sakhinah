package mocks

import (
	"clean_arc/features/users"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (m *Repository) InsertUser(newUser users.Core) (users.Core, error) {
	// ret := m.Called(newUser)

	// var r0 users.Core
	// if rf, ok := ret.Get(0).(func(users.Core) users.Core); ok {
	// 	r0 = rf(newUser)
	// } else {
	// 	r0 = ret.Get(0).(users.Core)
	// }

	// var r1 error
	// if rf, ok := ret.Get(1).(func(users.Core) error); ok {
	// 	r1 = rf(newUser)
	// } else {
	// 	r1 = ret.Error(1)
	// }

	// return r0, r1
	args := m.Called(newUser)
	result := args.Get(0).(users.Core)
	err := args.Error(1)
	return result, err
}

func (m *Repository) Login(email string, password string) (users.Core, error) {
	// ret := m.Called(email, password)

	// var r0 users.Core
	// if rf, ok := ret.Get(0).(func(string, string) users.Core); ok {
	// 	r0 = rf(email, password)
	// } else {
	// 	r0 = ret.Get(0).(users.Core)
	// }

	// var r1 error
	// if rf, ok := ret.Get(1).(func(string, string) error); ok {
	// 	r1 = rf(email, password)
	// } else {
	// 	r1 = ret.Error(1)
	// }

	// return r0, r1
	args := m.Called(email, password)
	result := args.Get(0).(users.Core)
	err := args.Error(1)
	return result, err
}

func (m *Repository) SelectAllUser() ([]users.CoreWithID, error) {
	args := m.Called()
	result := args.Get(0).([]users.CoreWithID)
	err := args.Error(1)
	return result, err
}

func (m *Repository) SelectUserByEmail(email string) (users.CoreWithID, error) {
	args := m.Called(email)
	result := args.Get(0).(users.CoreWithID)
	err := args.Error(1)
	return result, err
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
