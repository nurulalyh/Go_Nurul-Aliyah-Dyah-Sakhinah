package controllers

import (
	"net/http"
	"strconv"

	"ORM/eksplorasi/config"
	"ORM/eksplorasi/models"

	"github.com/labstack/echo"
)

func GetBlogsController(c echo.Context) error {
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
		"data":    responseBlogs,
	})
}

func GetBlogController(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	var blog []models.Blog
	if err := config.DB.First(&blog, "id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "blog not found",
		})
	}

	for i, data := range blog {
		var user models.User
		if err := config.DB.First(&user, data.UserID).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
				"message": "error get user data",
				"error":   err.Error(),
			})
		}
		blog[i].User = models.User{Name: user.Name}
	}

	var responseBlogs []map[string]interface{}
	for _, blog := range blog {
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
		"data":    responseBlogs,
	})
}

func CreateBlogController(c echo.Context) error {
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
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Id",
		})
	}

	var blog []models.Blog
	if err := config.DB.First(&blog, "id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "blog Not Found",
		})
	}

	if err := config.DB.Delete(&blog, "id = ?", id).Error; err != nil {
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
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Id",
		})
	}

	var blog models.Blog
	if err := config.DB.First(&blog, "id = ?", id).Error; err != nil {
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

	blog.Title = updateData.Title
    blog.Content = updateData.Content

	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to update user",
			"error":   err.Error(),
		})
	}

	var user models.User
    if err := config.DB.First(&user, blog.UserID).Error; err != nil {
        return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
            "message": "User Not Found",
        })
    }
    blog.User = user

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Success Update data",
        "data":    []models.Blog{blog},
    })
}
