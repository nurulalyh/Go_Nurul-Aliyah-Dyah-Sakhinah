package main

import (
	"fmt"
)

func kelipatan3(c chan int, n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			c <- i
		}
	}
	close(c)
}

func main() {
	var n int
	ch := make(chan int, 10)

	fmt.Print("Masukkan range bilangan untuk pencarian kelipatan 3: ")
	fmt.Scan(&n)
	fmt.Println()

	go kelipatan3(ch, n)

	for result := range ch {
		fmt.Println(result)
	}
}
