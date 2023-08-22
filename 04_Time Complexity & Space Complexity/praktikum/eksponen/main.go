package main

import (
	"fmt"
	"math"
)

func pow(x, n int) int {
	eksponen := math.Pow(float64(x), float64(n))
	return int(eksponen)
}

func pow2(x, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		res := pow2(x, n/2)
		return res * res
	} else {
		res := pow2(x, (n-1)/2)
		return x * res * res
	}
}

func main() {
	fmt.Println("Mencari Eksponen menggunakan function pow")
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
	fmt.Println("=====================")
	fmt.Println("Mencari Eksponen menggunakan function pow2")
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
