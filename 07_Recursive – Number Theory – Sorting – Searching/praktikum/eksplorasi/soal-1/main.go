package main


import "fmt"


func MaxSequence(arr []int) int {
	currMax := -10
	newMax := 0

	for _, v := range arr {
		newMax = newMax + v

		if newMax > currMax {
			currMax = newMax
		} else if newMax <= -1 {
			newMax = 0
		} 
	}
	return currMax
}


func main() {

fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) // 7

fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7

fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) // 8

fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) // 12

}

