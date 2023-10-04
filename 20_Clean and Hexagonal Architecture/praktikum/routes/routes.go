package routes

import (
	"clean_arc/features/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uh users.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/users/login", uh.Login())
	e.POST("/users", uh.CreateUser())
	e.GET("/users", uh.GetUsers())
}
