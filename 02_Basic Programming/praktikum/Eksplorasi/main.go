package main

import (
	"fmt"
	"strings"
)

func Palindrom(a string) bool {
	a = strings.ToLower(a)
	for i := 0; i < len(a)/2; i++ {
		if a[i] != a[len(a)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var kata string
	var tryAgain string
	
	for {
		fmt.Println("\t\tProgram Mengecek Palindrom 1 Kata")
		fmt.Println("----------------------------------------------------------------")
		
		fmt.Print("Masukkan satu kata tanpa spasi: ")
		fmt.Scanln(&kata)
		fmt.Println("----------------------------------------------------------------")
		
		
		if Palindrom(kata) {
			fmt.Printf("%s adalah palindrom.\n", kata)
		} else {
			fmt.Printf("%s bukan palindrom.\n", kata)
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