package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	mapCount := make(map[string]int)

	for _, item := range items {
		mapCount[item]++
	}

	var countRes []pair

	for k, v := range mapCount {
		countRes = append(countRes, pair{name: k, count: v})
	}

	sort.Slice(countRes, func(i, j int) bool {
		return countRes[i].count < countRes[j].count
	})

	return countRes
}

func main() {
	fmt.Println()
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()
	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()
	fmt.Println()
	
	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // football->1 basketball->1 tenis->1
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()
}
