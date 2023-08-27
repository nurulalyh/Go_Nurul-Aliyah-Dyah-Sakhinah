package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	calc := make(map[int]int)

	for _, str := range angka {
		num, _ := strconv.Atoi(string(str))
		calc[num]++
	}

	res := []int{}
	for num, count := range calc {
		if count == 1 {
			res = append(res, num)
		}
	}

	return res
}

func main() {
	// Test cases
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
