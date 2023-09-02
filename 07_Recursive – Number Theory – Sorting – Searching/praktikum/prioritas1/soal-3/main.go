package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}

	akarKuadrat := math.Sqrt(float64(number))

	for i := 2; i <= int(akarKuadrat); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func primeX(number int) int {
	index := 0
	checkPrime := 1

	for index < number {
        checkPrime++
		if primeNumber(checkPrime) {
			index++
		}
	}

	return checkPrime
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
