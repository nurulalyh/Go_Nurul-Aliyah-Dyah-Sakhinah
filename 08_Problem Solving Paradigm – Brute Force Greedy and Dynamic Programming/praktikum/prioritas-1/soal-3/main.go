package main

import "fmt"

var (
	tabelFibo = make(map[int]int)
)

func fibonacci(number int) int {
	if number <= 1 {
		tabelFibo[number] = number
		return number
	}

	if k, v := tabelFibo[number]; v {
		return k
	}

	tabelFibo[number] = fibonacci(number-2) + fibonacci(number-1)

	return fibonacci(number-2) + fibonacci(number-1)
}

func main() {
	fmt.Println(fibonacci(0)) // 0

	fmt.Println(fibonacci(2)) // 1

	fmt.Println(fibonacci(9)) // 34

	fmt.Println(fibonacci(10)) // 55

	fmt.Println(fibonacci(12)) // 144
}