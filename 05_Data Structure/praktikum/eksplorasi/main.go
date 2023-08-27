package main

import (
	"fmt"
	"math"
)

func selisihAbs(matriks [][]int) int {
	var size int = len(matriks)
	var diagonalKiri int = 0
	var diagonalKanan int = 0

	for i := 0; i < size; i++ {
		diagonalKiri += matriks[i][i]
		diagonalKanan += matriks[i][size-i-1]
	}

	return int(math.Abs(float64(diagonalKiri - diagonalKanan)))
}

func inputMatriks() [][]int {
	var size int
	fmt.Print("Masukkan ukuran matriks : ")
	fmt.Scan(&size)

	matriks := make([][]int, size)
	for i := 0; i < size; i++ {
		matriks[i] = make([]int, size)
		for j := 0; j < size; j++ {
			fmt.Printf("Masukkan nilai untuk indeks (%d, %d): ", i, j)
			fmt.Scan(&matriks[i][j])
		}
	}
	fmt.Println("\nMatriksnya : ")
	fmt.Println(matriks)
	return matriks
}

func main() {
	var hasil int = int(selisihAbs(inputMatriks()))
	fmt.Println("\nselisih absolut antara jumlah diagonalnya adalah ", hasil)
}
