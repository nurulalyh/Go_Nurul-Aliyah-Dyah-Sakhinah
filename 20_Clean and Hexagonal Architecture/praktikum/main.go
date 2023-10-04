package main

import (
	"clean_arc/config"
	"clean_arc/features/users/handler"
	"clean_arc/features/users/repository"
	"clean_arc/features/users/usecase"
	"clean_arc/routes"
	"clean_arc/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	dbConn := database.InitDB(*cfg)
	database.Migrate(dbConn)

	repo := repository.NewUserModel(dbConn)
	srv := usecase.NewUserLogic(repo)
	hdl := handler.NewUserController(srv)

	routes.InitRoute(e, hdl)

	if err := e.Start(":8000"); err != nil {
		e.Logger.Fatal("cannot start server ", err.Error())
	}
}
