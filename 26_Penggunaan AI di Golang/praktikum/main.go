package main

import (
	"fmt"
	"golang_ai/handler"
	"golang_ai/usecase"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load .env file. Err: %s", err)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	laptopUsecase := usecase.NewLaptopUsecase()
	laptopHandler := handler.NewLaptopHandler(laptopUsecase)
	e.POST("/laptop", laptopHandler.RecommendLaptop)
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	e.Start(port)
}
