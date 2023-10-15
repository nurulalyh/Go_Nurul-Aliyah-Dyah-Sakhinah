package handler

import (
	"fmt"
	"golang_ai/usecase"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type LaptopHandler struct {
	LaptopUsecase usecase.LaptopUsecase // Inject the use case
}

type LaptopResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func NewLaptopHandler(usecase usecase.LaptopUsecase) *LaptopHandler {
	return &LaptopHandler{
		LaptopUsecase: usecase,
	}
}

func (h *LaptopHandler) RecommendLaptop(c echo.Context) error {
	var requestData map[string]interface{}
	err := c.Bind(&requestData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON format")
	}

	budget, okBudget := requestData["budget"].(float64)
	purpose, okPurpose := requestData["purpose"].(string)
	brand, okBrand := requestData["brand"].(string)
	ram, okRAM := requestData["ram"].(string)
	cpu, okCPU := requestData["cpu"].(string)
	screenSize, okScreenSize := requestData["screen_size"].(string)

	if !okBudget || !okPurpose || !okBrand || !okRAM || !okCPU || !okScreenSize {
		return c.JSON(http.StatusBadRequest, "Invalid request format")
	}

	userInput := fmt.Sprintf("Rekomendasi Laptop merk %s untuk %s dengan spesifikasi %s ram %s cpu %s ukuran layar dan dengan budget sebesar Rp %.0f.", brand, purpose, ram, cpu, screenSize, budget)

	answer, err := h.LaptopUsecase.RecommendLaptop(userInput, brand, ram, cpu, screenSize, os.Getenv("API_KEY"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error dalam mmebuat rekomendasi")
	}

	responseData := LaptopResponse{
		Status: "success",
		Data:   answer,
	}

	return c.JSON(http.StatusOK, responseData)
}
