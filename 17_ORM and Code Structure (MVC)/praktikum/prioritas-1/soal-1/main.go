package main

import (
	"fmt"
	"net/http"
  	"strconv"

  	"github.com/jinzhu/gorm"
  	_ "github.com/jinzhu/gorm/dialects/mysql"
  	"github.com/labstack/echo"
)

var (
	DB *gorm.DB
)

func init() {
  	InitDB()
  	InitialMigration()
}

type Config struct {
  	DB_Username string
  	DB_Password string
  	DB_Port     string
  	DB_Host     string
  	DB_Name     string
}

func InitDB() {
	config := Config{
    	DB_Username: "root",
    	DB_Password: "",
    	DB_Port:     "3306",
    	DB_Host:     "localhost",
    	DB_Name:     "crud_go",
  	}

  	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
    	config.DB_Username,
    	config.DB_Password,
    	config.DB_Host,
    	config.DB_Port,
    	config.DB_Name,
  	)

  	var err error
  	DB, err = gorm.Open("mysql", connectionString)
  	if err != nil {
    	panic(err)
  	}
}

type User struct {
  	gorm.Model
  	Name     string `json:"name" form:"name"`
  	Email    string `json:"email" form:"email"`
  	Password string `json:"password" form:"password"`
}

func InitialMigration() {
  	DB.AutoMigrate(&User{})
}

func GetUsersController(c echo.Context) error {
  	var users []User

  	if err := DB.Find(&users).Error; err != nil {
    	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  	}

  	return c.JSON(http.StatusOK, map[string]interface{}{
    	"message": "success get all users",
    	"users":   users,
  	})
}

func GetUserController(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Id",
		})
	}

	var user []User
	if err := DB.First(&user, id).Error; err != nil {
    	return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "User Not Found",
		})
  	}

	return c.JSON(http.StatusOK, map[string]interface{}{
    	"message": "success get user",
    	"users":   user,
  	})
}

func CreateUserController(c echo.Context) error {
  	user := User{}

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

  	if err := DB.Create(&user).Error; err != nil {
    	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  	}

  	return c.JSON(http.StatusOK, map[string]interface{}{
    	"message": "success create new user",
    	"user":    user,
  	})
}

func DeleteUserController(c echo.Context) error {
  	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Id",
		})
	}

	var user User
	if err := DB.First(&user, id).Error; err != nil {
    	return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "User Not Found",
		})
  	}

	if err := DB.Delete(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
    	"message": "success delete data",
  	})
}

func UpdateUserController(c echo.Context) error {
  	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Id",
		})
	}

	user := User{}
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var existingUser User
	if err := DB.First(&existingUser, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "User Not Found",
		})
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := DB.Save(&existingUser).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Update data",
		"user":    existingUser,
	})
}

func main() {
  	e := echo.New()
  
  	e.GET("/users", GetUsersController)
  	e.GET("/users/:id", GetUserController)
  	e.POST("/users", CreateUserController)
  	e.DELETE("/users/:id", DeleteUserController)
  	e.PUT("/users/:id", UpdateUserController)

  	e.Logger.Fatal(e.Start(":8000"))
}