package main

import (
	"fmt"
)

func romawi(num int) string {
	romawiArr := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	numberArr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""

	for i := 0; i < len(romawiArr); i++ {
		for num >= numberArr[i] {
			result += romawiArr[i]
			num -= numberArr[i]
		}
	}

	return result
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&n)
	fmt.Println()

	fmt.Println(romawi(n))
}
