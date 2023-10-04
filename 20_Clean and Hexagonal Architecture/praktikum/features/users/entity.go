package users

import (
	
	"github.com/labstack/echo/v4"
)

type Core struct {
	Email    string
	Password string
}

type CoreWithID struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Handler interface { //dibutuh newuserroute
	CreateUser() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetUsers() echo.HandlerFunc
}

type UseCase interface { //dibutuh newusercontroller
	CreateUser(newUser Core) error
	Login(email string, password string) (Core, error)
	GetUsers() ([]CoreWithID, error)
	GetUserByEmail(email string) (CoreWithID, error)
}

type Repository interface { //dibutuh newuserlogic
	InsertUser(newUser Core) (Core, error)
	Login(email string, password string) (Core, error)
	SelectAllUser() ([]CoreWithID, error)
	SelectUserByEmail(email string) (CoreWithID, error)
}
