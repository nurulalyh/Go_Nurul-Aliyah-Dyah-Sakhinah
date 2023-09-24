package main

import (
	"ORM/prioritas-1/soal-2/config"
	"ORM/prioritas-1/soal-2/routes"

	"github.com/labstack/echo"
)

func main() {
	config.InitDB()
	e := echo.New()

	routes.SetUserRoutes(e)
	
	e.Logger.Fatal(e.Start(":8000"))
}
