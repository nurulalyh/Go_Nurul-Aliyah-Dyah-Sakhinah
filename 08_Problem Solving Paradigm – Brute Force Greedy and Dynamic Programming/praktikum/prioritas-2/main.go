package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)

	if n <= 1 {
		return 0
	}

	minCost := make([]int, n)
	minCost[0] = 0
	minCost[1] = int(math.Abs(float64(jumps[1] - jumps[0])))

	for i := 2; i < n; i++ {
		cost1 := minCost[i-1] + int(math.Abs(float64(jumps[i]-jumps[i-1])))
		cost2 := minCost[i-2] + int(math.Abs(float64(jumps[i]-jumps[i-2])))
		minCost[i] = int(math.Min(float64(cost1), float64(cost2)))
	}

	return minCost[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}