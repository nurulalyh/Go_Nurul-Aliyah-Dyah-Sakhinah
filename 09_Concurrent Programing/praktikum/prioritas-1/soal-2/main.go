package main

import (
	"fmt"
	"time"
)

func kelipatan3(c chan int, n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			c <- i
		}
	}
}

func main() {
	var n int
	ch := make(chan int)

	fmt.Print("Masukkan range bilangan untuk pencarian kelipatan 3: ")
	fmt.Scan(&n)
	fmt.Println()

	go kelipatan3(ch, n)

	go func()  {
		time.Sleep(1 * time.Second)
		close(ch)
	}()
	
	for result := range ch {
		fmt.Println(result)
	}
}