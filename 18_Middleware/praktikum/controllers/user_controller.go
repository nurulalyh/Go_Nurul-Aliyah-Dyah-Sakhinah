package controllers

import (
	"middleware/config"
	"middleware/helper"
	"middleware/models"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "need authorization token")
	}

	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return helper.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization token")
	}

	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"user":    users,
	})
}

func GetUserController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "need authorization token")
	}

	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return helper.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error in token claims")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token does not match user ID")
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Id",
		})
	}

	if int(userID) != id {
		return echo.NewHTTPError(http.StatusForbidden, "You do not have access to this data")
	}

	var user []models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	var user models.User

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
		"message": "success create new user, please login to get token",
		"user":    user,
	})
}

func LoginUserController(c echo.Context) error {
	request := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user models.User
	if err := config.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, " email or password does not match")
	}

	if user.Password != request.Password {
		return echo.NewHTTPError(http.StatusUnauthorized, "email or password does not match")
	}

	token, err := helper.CreateJWTToken(user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed in creating JWT token")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login",
		"user_id": user.ID,
		"token":   token,
	})
}

func DeleteUserController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "need authorization token")
	}

	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return helper.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error in token claims")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token does not match user ID")
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Id",
		})
	}

	if int(userID) != id {
		return echo.NewHTTPError(http.StatusForbidden, "You do not have access to this data")
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "User Not Found",
		})
	}

	if err := config.DB.Delete(&user, id).Error; err != nil {
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
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "need authorization token")
	}

	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return helper.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error in token claims")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token does not match user ID")
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Id",
		})
	}

	if int(userID) != id {
		return echo.NewHTTPError(http.StatusForbidden, "You do not have access to this data")
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
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
		"user":    user,
	})
}
