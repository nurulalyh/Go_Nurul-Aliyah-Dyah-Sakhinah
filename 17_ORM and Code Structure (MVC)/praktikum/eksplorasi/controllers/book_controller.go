package controllers

import (
	"net/http"
	"strconv"

	"ORM/prioritas-2/config"
	"ORM/prioritas-2/models"

	"github.com/labstack/echo"
)

func GetBooksController(c echo.Context) error {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"data":    books,
	})
}

func GetBookController(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	var book []models.Book
	if err := config.DB.First(&book, "id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "book not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"data":    book,
	})
}

func CreateBookController(c echo.Context) error {
	book := models.Book{}

	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Error when parsing data",
			"error":   err.Error(),
		})
	}

	if err := config.DB.Create(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"data":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Id",
		})
	}

	var book []models.Book
	if err := config.DB.First(&book, "id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "book Not Found",
		})
	}

	if err := config.DB.Delete(&book, "id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "error delete user",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete data",
	})
}

func UpdateBookController(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Id",
		})
	}

	var book models.Book
	if err := config.DB.First(&book, "id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "book Not Found",
		})
	}

	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Error when parsing data",
			"error":   err.Error(),
		})
	}

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to update user",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Update data",
		"data":    book,
	})
}
