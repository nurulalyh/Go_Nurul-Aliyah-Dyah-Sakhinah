package usecase

import (
	"clean_arc/features/users"
	"errors"
	"strings"

	"github.com/sirupsen/logrus"
)

type userLogic struct {
	m users.Repository
}

func NewUserLogic(r users.Repository) users.UseCase {
	return &userLogic{
		m: r,
	}
}

func (ul *userLogic) CreateUser(newUser users.Core) error {
	_, err := ul.m.InsertUser(newUser)
	if err != nil {
		logrus.Error("create user logic error:", err.Error())
		return errors.New("error from server")
	}

	return nil
}

func (ul *userLogic) Login(email string, password string) (users.Core, error) {
	result, err := ul.m.Login(email, password)
	if err != nil {
		if strings.Contains(err.Error(), "appropriate") || strings.Contains(err.Error(), "contain") {
			return users.Core{}, err
		}
		return users.Core{}, errors.New("error from server")
	}

	return result, nil
}

func (ul *userLogic) GetUsers() ([]users.CoreWithID, error) {
	usr, err := ul.m.SelectAllUser()
	if err != nil {
		return []users.CoreWithID{}, errors.New("error from server")
	}

	return usr, nil
}

func (ul *userLogic) GetUserByEmail(email string) (users.CoreWithID, error) {
	user, err := ul.m.SelectUserByEmail(email)
	if err != nil {
		return users.CoreWithID{}, errors.New("error from server")
	}

	return user, nil
}
