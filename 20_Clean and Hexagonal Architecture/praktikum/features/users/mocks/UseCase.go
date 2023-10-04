package mocks

import (
	"clean_arc/features/users"

	"github.com/stretchr/testify/mock"
)

type UseCase struct {
	mock.Mock
}

func (_m *UseCase) Login(hp string, password string) (users.Core, error) {
	ret := _m.Called(hp, password)

	var r0 users.Core
	if rf, ok := ret.Get(0).(func(string, string) users.Core); ok {
		r0 = rf(hp, password)
	} else {
		r0 = ret.Get(0).(users.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(hp, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UseCase) Register(newUser users.Core) error {
	ret := _m.Called(newUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUseCase interface {
	mock.TestingT
	Cleanup(func())
}

func NewUseCase(t mockConstructorTestingTNewUseCase) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
