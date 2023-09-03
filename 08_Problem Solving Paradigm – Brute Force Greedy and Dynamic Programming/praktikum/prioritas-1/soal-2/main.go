package main

import "fmt"

func pascalTriangle(numRows int) [][]int {
    if numRows <= 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}

	segitiga := pascalTriangle(numRows - 1)
	row := segitiga[len(segitiga)-1]
	new := make([]int, numRows)
	new[0], new[numRows - 1] = 1, 1

	for i := 1; i < numRows-1; i++ {
		new[i] = row[i-1] + row[i]
	}

	segitiga = append(segitiga, new)
	return segitiga
}

func main() {
    var numRows int

	fmt.Print("Masukkan Jumlah Row Segitiga Pascal : ")
	fmt.Scan(&numRows)

    fmt.Println(pascalTriangle(numRows))
}
