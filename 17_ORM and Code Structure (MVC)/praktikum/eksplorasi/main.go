package main

import (
	"ORM/eksplorasi/config"
	"ORM/eksplorasi/routes"

	"github.com/labstack/echo"
)

func main() {
	config.InitDB()
	e := echo.New()

	routes.SetUserRoutes(e)
	routes.SetBookRoutes(e)
	routes.SetBlogRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
