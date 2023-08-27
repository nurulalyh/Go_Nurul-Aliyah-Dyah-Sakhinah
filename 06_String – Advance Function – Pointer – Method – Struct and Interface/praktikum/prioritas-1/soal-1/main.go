package main

import "fmt"

type Car struct {
	carType string
	fuelin  float64
}

func (c Car) mileage() float64 {
	var fuelConsume float64 = 1.5
	var estimateMile float64 = c.fuelin * fuelConsume
	return estimateMile
}

func main() {
	var Car1 = Car{
		carType: "SUV",
		fuelin:  10.5,
	}

	var estimateMile float64 = Car1.mileage()
	fmt.Println("Car type: ", Car1.carType)
	fmt.Println("Fuelin: ", Car1.fuelin)
	fmt.Println("Estimate Distance (Mile): ", estimateMile)
}