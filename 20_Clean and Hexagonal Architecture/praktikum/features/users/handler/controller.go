package handler

import (
	"clean_arc/features/users"
	"clean_arc/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type userController struct {
	service users.UseCase
}

func NewUserController(us users.UseCase) users.Handler {
	return &userController{
		service: us,
	}
}

func (uc *userController) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := CreateUserInput{}
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Error Bind", err.Error())
			return c.JSON(helper.ResponsFormat(http.StatusBadRequest, "Input error from user", nil))
		}

		err := uc.service.CreateUser(users.Core{Email: input.Email, Password: input.Password})
		if err != nil {
			return c.JSON(helper.ResponsFormat(http.StatusInternalServerError, err.Error(), nil))
		}

		user, err := uc.service.GetUserByEmail(input.Email)
		if err != nil {
			return c.JSON(helper.ResponsFormat(http.StatusInternalServerError, err.Error(), nil))
		}

		return c.JSON(helper.ResponsFormat(http.StatusCreated, "Success create data", user))
	}
}

func (uc *userController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginInput
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Error bind, ", err.Error())
			return c.JSON(helper.ResponsFormat(http.StatusBadRequest, "Error input from user", nil))
		}

		user, err := uc.service.GetUserByEmail(input.Email)
		if err != nil {
			code := http.StatusInternalServerError
			if strings.Contains(err.Error(), "appropriate") || strings.Contains(err.Error(), "contain") {
				code = http.StatusBadRequest
			}
			return c.JSON(helper.ResponsFormat(code, "Error from server or user", nil))
		}

		res, err := uc.service.Login(input.Email, input.Password)
		if err != nil {
			code := http.StatusInternalServerError
			if strings.Contains(err.Error(), "appropriate") || strings.Contains(err.Error(), "contain") {
				code = http.StatusBadRequest
			}
			return c.JSON(helper.ResponsFormat(code, "Error from server or user", nil))
		}

		var result = new(LoginResponse)
		token := helper.GenerateJWT(strconv.Itoa(int(user.Id)))
		result.Email = res.Email
		result.Token = token

		return c.JSON(helper.ResponsFormat(http.StatusOK, "Success login", result))
	}
}

func (uc *userController) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.service.GetUsers()
		if err != nil {
			return c.JSON(helper.ResponsFormat(http.StatusInternalServerError, err.Error(), nil))
		}

		return c.JSON(helper.ResponsFormat(http.StatusOK, "Success get all users", users))
	}
}
