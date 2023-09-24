package routes

import (
	"ORM/prioritas-2/controllers"
	"github.com/labstack/echo"
)

func SetBookRoutes(e *echo.Echo) {
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
}
