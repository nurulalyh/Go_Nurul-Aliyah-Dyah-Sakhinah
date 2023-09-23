package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success get all users",
		"users":    users,
	})
}

func GetUserController(c echo.Context) error {
	var index = c.Param("id")
	idx, err := strconv.Atoi(index)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid User Id",
		})
	}

	if idx < 0 || idx > len(users)-1 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data Not Found",
		}) 
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data": users[idx-1],
	})
}

func DeleteUserController(c echo.Context) error {
	var index = c.Param("id")
	idx, err := strconv.Atoi(index)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid User Id",
		})
	}

	for i, user := range users {
		if user.Id == idx {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "data deleted successfully",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "Data Not Found",
	})
}

func UpdateUserController(c echo.Context) error {
	var index = c.Param("id")
	idx, err := strconv.Atoi(index)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
			"error": err.Error(),
		})
	}

	for i, user := range users {
		if user.Id == idx {
			c.Bind(&user)
			users[i] = user
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "data updated successfully",
				"data": user,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "Data Not Found",
		"error": err.Error(),
	})
}

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	e := echo.New()

	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	e.Logger.Fatal(e.Start(":8000").Error())
}
