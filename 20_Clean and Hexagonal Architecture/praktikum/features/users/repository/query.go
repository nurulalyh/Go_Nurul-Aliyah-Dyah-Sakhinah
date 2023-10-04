package repository

import (
	"clean_arc/features/users"
	"clean_arc/helper"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type usersModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) users.Repository {
	return &usersModel{
		db: db,
	}
}

func (um *usersModel) InsertUser(newUser users.Core) (users.Core, error) {
	var inputUser = User{}

	hashedPassword, err := helper.GenerateHashedPassword(newUser.Password)
	if err != nil {
		logrus.Error("Error when hashing password ", err.Error())
		return users.Core{}, err
	}

	inputUser.Email = newUser.Email
	inputUser.Password = hashedPassword
	if err := um.db.Table("users").Create(&inputUser).Error; err != nil {
		logrus.Error("Model : Create data error, ", err.Error())
		return users.Core{}, err
	}

	return newUser, nil
}

func (um *usersModel) Login(email, password string) (users.Core, error) {
	res := User{}
	if err := um.db.Where("email = ? ", email).Find(&res).Error; err != nil {
		logrus.Error("Login error, ", err.Error())
		return users.Core{}, err
	}

	if res.Email == "" {
		return users.Core{}, errors.New("data not found")
	}

	if !helper.ComparePassword(res.Password, password) {
		return users.Core{}, errors.New("password tidak sesuai")
	}

	return users.Core{Email: res.Email}, nil
}

func (um *usersModel) SelectAllUser() ([]users.CoreWithID, error) {
	usr := []User{}

	if err := um.db.Find("&users").Error; err != nil {
		logrus.Error("Model : Select data error, ", err.Error())
		return []users.CoreWithID{}, err
	}

	var coreWithIDs []users.CoreWithID

	for _, user := range usr {
		coreWithID := users.CoreWithID{}
		coreWithID.Id = user.ID
		coreWithID.Email = user.Email
		coreWithID.Password = user.Password

		coreWithIDs = append(coreWithIDs, coreWithID)
	}

	return []users.CoreWithID{}, nil
}

func (um *usersModel) SelectUserByEmail(email string) (users.CoreWithID, error) {
	var user users.CoreWithID

	if err := um.db.Where("email = ?", email).First(&user).Error; err != nil {
		return users.CoreWithID{}, err
	}

	return user, nil
}
