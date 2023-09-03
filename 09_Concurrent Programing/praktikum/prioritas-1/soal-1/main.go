package main

import (
	"fmt"
	"time"
)

func kelipatan(x int, n int) {
	for i := 1; i <= n; i++ {
		if i%x == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	var x, n int

	fmt.Print("Masukkan angka yang ingin dicari kelipatannya : ")
	fmt.Scan(&x)
	fmt.Println()
	fmt.Print("Masukkan range bilangan pencarian kelipatan : ")
	fmt.Scan(&n)
	fmt.Println()

	go kelipatan(x,n)
	time.Sleep(3 * time.Second)
}