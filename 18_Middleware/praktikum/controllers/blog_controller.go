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

func GetBlogsController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token dibutuhkan")
	}
	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return helper.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token salah")
	}

	var blogs []models.Blog

	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	for i, blog := range blogs {
		var user models.User
		if err := config.DB.First(&user, blog.UserID).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
				"message": "error get user data",
				"error":   err.Error(),
			})
		}
		blogs[i].User = models.User{Name: user.Name}
	}

	var responseBlogs []map[string]interface{}
	for _, blog := range blogs {
		user := map[string]interface{}{
			"name": blog.User.Name,
		}

		responseBlogs = append(responseBlogs, map[string]interface{}{
			"ID":      blog.ID,
			"Title":   blog.Title,
			"Content": blog.Content,
			"User":    user,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   responseBlogs,
	})
}

func GetBlogController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token dibutuhkan")
	}
	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return helper.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token salah")
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	blog := models.Blog{}
	if err := config.DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "blog not found",
		})
	}

	if err := c.Bind(&blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Error when parsing data",
			"error":   err.Error(),
		})
	}

	var user models.User
	if err := config.DB.First(&user, blog.UserID).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": "error get blog data",
			"error":   err.Error(),
		})
	}

	userName := user.Name
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog": map[string]interface{}{
			"title":   blog.Title,
			"content": blog.Content,
			"user": map[string]interface{}{
				"name": userName,
			},
		},
	})
}

func CreateBlogController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token dibutuhkan")
	}
	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return helper.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token salah")
	}

	blog := models.Blog{}

	if err := c.Bind(&blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Error when parsing data",
			"error":   err.Error(),
		})
	}

	var user models.User
	if err := config.DB.First(&user, blog.UserID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	if err := config.DB.Create(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	userName := user.Name
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog": map[string]interface{}{
			"title":   blog.Title,
			"content": blog.Content,
			"user": map[string]interface{}{
				"name": userName,
			},
		},
	})
}

func DeleteBlogController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token dibutuhkan")
	}
	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return helper.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token salah")
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Id",
		})
	}

	var blog []models.Blog
	if err := config.DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "blog Not Found",
		})
	}

	if err := config.DB.Delete(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "error delete user",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete data",
	})
}

func UpdateBlogController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token dibutuhkan")
	}
	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return helper.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token salah")
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Id",
		})
	}

	var blog models.Blog
	if err := config.DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "blog Not Found",
		})
	}

	var updateData models.Blog
	if err := c.Bind(&updateData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Error when parsing data",
			"error":   err.Error(),
		})
	}

	blog.UserID = updateData.UserID
	blog.Title = updateData.Title
	blog.Content = updateData.Content

	var user models.User
	if err := config.DB.First(&user, blog.UserID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "User Not Found",
		})
	}

	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to update blog",
			"error":   err.Error(),
		})
	}

	userName := user.Name
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
		"blog": map[string]interface{}{
			"title":   blog.Title,
			"content": blog.Content,
			"user": map[string]interface{}{
				"name": userName,
			},
		},
	})
}
