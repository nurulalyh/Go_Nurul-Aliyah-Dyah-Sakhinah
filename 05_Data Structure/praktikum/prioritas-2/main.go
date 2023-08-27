
package main

import "fmt"

func PairSum(arr []int, target int) []int {
	var indexLeft int = 0
	var indexRight int = len(arr) - 1

	for indexLeft < indexRight {
		tot := arr[indexLeft] + arr[indexRight]

		if tot == target {
			return []int{indexLeft, indexRight}
		} else if tot < target {
			indexLeft++
		} else {
			indexRight--
		}
	}

	return []int{-1,-1}
}

func main() {
	// Test cases
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}