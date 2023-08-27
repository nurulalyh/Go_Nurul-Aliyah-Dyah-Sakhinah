package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min, max int) {
	min = *numbers[0]
	max = *numbers[0]

	for i, _ := range numbers {
		if *numbers[i] < min {
			min = *numbers[i]
		}
		if *numbers[i] > max {
			max = *numbers[i]
		}
	}
	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Println("Input: ")
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	fmt.Println()

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Output: ")
	fmt.Println(max, "is the maximum number")
	fmt.Println(min, "is the minimum number")
}