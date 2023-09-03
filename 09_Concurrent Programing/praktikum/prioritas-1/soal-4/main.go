package main

import (
	"fmt"
	"sync"
)

func faktorial(n int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	fmt.Print("Faktorial dari ", n, " adalah ", result)
	mu.Unlock()
}

func main() {
	var num int

	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&num)

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go faktorial(num, &wg, &mu)
	wg.Wait()
}
