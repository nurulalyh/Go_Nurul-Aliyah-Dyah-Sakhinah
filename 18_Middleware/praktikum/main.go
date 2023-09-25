package main

import (
	"middleware/config"
	"middleware/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config.InitDB()
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.SetUserRoutes(e)
	routes.SetBookRoutes(e)
	routes.SetBlogRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
