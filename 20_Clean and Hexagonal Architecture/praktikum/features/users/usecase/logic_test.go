package usecase_test

import (
	"clean_arc/features/users"
	"clean_arc/features/users/mocks"
	"clean_arc/features/users/usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	repo := mocks.NewRepository(t)
	ul := usecase.NewUserLogic(repo)
	successCaseData := users.Core{Email: "nurulaliyah2303@gmail.com", Password: "morilla_arta"}

	t.Run("Success login", func(t *testing.T) {
		repo.On("Login", successCaseData.Email, successCaseData.Password).Return(users.Core{Email: "nurulaliyah2303@gmail.com"}, nil).Once()
		result, err := ul.Login("nurulaliyah2303@gmail.com", "morilla_arta")

		assert.Nil(t, err)
		assert.Equal(t, "nurulaliyah2303@gmail.com", result.Email)
		repo.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		repo.On("Login", successCaseData.Email, "morilla_lmj").Return(users.Core{}, errors.New("wrong password")).Once()
		result, err := ul.Login("nurulaliyah2303@gmail.com", "morilla_lmj")

		assert.Error(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Empty(t, result.Email)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("Login", "nurul@gmail.com", "morilla_lmj").Return(users.Core{}, errors.New("Data not found")).Once()
		result, err := ul.Login("nurul@gmail.com", "morilla_lmj")

		assert.Error(t, err)
		assert.ErrorContains(t, err, "error from server")
		assert.Empty(t, result.Email)
		repo.AssertExpectations(t)
	})

	t.Run("Error from server", func(t *testing.T) {
		repo.On("Login", successCaseData.Email, "morilla_lmj").Return(users.Core{}, errors.New("column not exist")).Once()
		result, err := ul.Login("nurulaliyah2303@gmail.com", "morilla_lmj")

		assert.Error(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Empty(t, "", result.Email)
		repo.AssertExpectations(t)
	})
}

func TestCreateUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	ul := usecase.NewUserLogic(repo)
	insertData := users.Core{
		Email:    "nurulaliyah2303@gmail.com",
		Password: "morilla_arta",
	}

	t.Run("Seccess create user", func(t *testing.T) {
		repo.On("InsertUser", insertData).Return(insertData, nil).Once()
		err := ul.CreateUser(insertData)

		assert.Nil(t, err)
	})

	t.Run("Failed create user", func(t *testing.T) {
		repo.On("InsertUser", insertData).Return(users.Core{}, errors.New("too many values")).Once()
		err := ul.CreateUser(insertData)

		assert.Error(t, err)
	})
}
