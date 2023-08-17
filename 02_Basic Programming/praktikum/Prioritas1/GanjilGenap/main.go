package main

import "fmt"

func main() {
	var bil int8
	var tryAgain string

	for {
		fmt.Println("\tProgram Menentukan Bilangan Ganjil atau Genap")
		fmt.Println("-------------------------------------------------------------")
		fmt.Print("Masukkan Bilangan Bulat: ")
		fmt.Scanln(&bil)
		fmt.Println("")
		fmt.Println("-------------------------------------------------------------")

		if bil%2 == 0 {
			fmt.Printf("%v adalah bilangan genap\n", bil)
			fmt.Println("-------------------------------------------------------------")
		} else {
			fmt.Printf("%v adalah bilangan ganjil\n", bil)
			fmt.Println("-------------------------------------------------------------")
		}

		fmt.Print("Apakah Anda ingin mencoba lagi? (y/n): ")
		fmt.Scanln(&tryAgain)
		fmt.Println("-------------------------------------------------------------")
		if tryAgain != "y" && tryAgain != "Y" {
			fmt.Println("Terima Kasih Telah Menggunakan Program Ini 😊")
			fmt.Println("-------------------------------------------------------------")
			break
		}
	}
}
