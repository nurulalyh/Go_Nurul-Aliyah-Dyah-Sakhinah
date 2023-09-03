package main

import "fmt"

func binary(index int) []string {
	binarySlice := make([]string, 0)

	for i := 0; i < (index + 1); i++ {
		binaryStr := fmt.Sprintf("%b", i)
		binarySlice = append(binarySlice, binaryStr)
	}

	return binarySlice
}

func main() {
	var n int

	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	fmt.Println(binary(n))
}
