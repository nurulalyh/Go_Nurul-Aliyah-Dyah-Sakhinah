package main

import (
	"fmt"
	"strings"
	"sync"
)

func charFrequency(sentence string, frequency chan map[rune]int, wg *sync.WaitGroup) {
	defer wg.Done()
	countMap := make(map[rune]int)

	for _, char := range sentence {
		if char != ' ' && (char < 'A' || (char > 'Z' && char < 'a') || char > 'z') {
			continue
		}
		countMap[char]++
	}

	frequency <- countMap
}

func main() {
	result := make(chan map[rune]int)
	numGoroutines := 4
	var wg sync.WaitGroup
	
	sentence := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	
	split := strings.Split(sentence, " ")

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go charFrequency(split[i], result, &wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	fixResult := make(map[rune]int)

	for letterCount := range result {
		for char, count := range letterCount {
			fixResult[char] += count
		}
	}

	for char, count := range fixResult {
		fmt.Printf("%c: %d\n", char, count)
	}
}