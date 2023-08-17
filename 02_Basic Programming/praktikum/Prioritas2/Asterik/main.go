package main

import (
	"fmt"
	"strings"
)

func main() {
	var tinggi int
	var tryAgain string

	for {
		fmt.Println("\t\tProgram Mencetak Segitiga Asterik")
		fmt.Println("----------------------------------------------------------------")
		
		fmt.Print("Masukkan tinggi segitiga: ")
		fmt.Scanln(&tinggi)
		fmt.Println("----------------------------------------------------------------")
		
		for i := 1; i <= tinggi; i++ {
			spasi := strings.Repeat(" ", tinggi-i)
			segitiga := strings.Repeat("* ", i)
			fmt.Print(spasi, segitiga, "\n")
		}
		fmt.Println("----------------------------------------------------------------")

		fmt.Print("Apakah Anda ingin mencoba lagi? (y/n): ")
		fmt.Scanln(&tryAgain)
		fmt.Println("----------------------------------------------------------------")
		if tryAgain != "y" && tryAgain != "Y" {
			fmt.Println("Terima Kasih Telah Menggunakan Program Ini 😊")
			fmt.Println("----------------------------------------------------------------")
			break
		}
	}
}
