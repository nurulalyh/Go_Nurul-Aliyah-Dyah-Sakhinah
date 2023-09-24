package controllers

import (
	"net/http"
	"strconv"

	"ORM/prioritas-2/config"
	"ORM/prioritas-2/models"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"data":    users,
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

	var user []models.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"data":    user,
	})
}

func CreateUserController(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Error when parsing data",
			"error":   err.Error(),
		})
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"data":    user,
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

	var user []models.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "User Not Found",
		})
	}

	if err := config.DB.Delete(&user, "id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "error delete user",
			"error":   err.Error(),
		})
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

	var user models.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "User Not Found",
		})
	}

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Error when parsing data",
			"error":   err.Error(),
		})
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to update user",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Update data",
		"data":    user,
	})
}
