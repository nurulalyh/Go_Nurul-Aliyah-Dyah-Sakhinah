package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	mergedArray := append(arrayA, arrayB...)
	mapping := make(map[string]bool)
	newArray := make([]string, 0)

	for _, value := range mergedArray {
		if !mapping[value] {
			mapping[value] = true
			newArray = append(newArray, value)
		}
	}
	return newArray
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}