package main

import "fmt"

func main() {
    var bil int
	var tryAgain string

	for {
		fmt.Println("\t\tProgram Mencetak Faktor Bilangan")
		fmt.Println("----------------------------------------------------------------")
		
		fmt.Print("Masukkan bil: ")
		fmt.Scanln(&bil)
		fmt.Println("----------------------------------------------------------------")
		
		fmt.Printf("Faktor dari %d adalah: \n", bil)
		for i := 1; i <= bil; i++ {
			if bil%i == 0 {
				fmt.Printf("%d\n", i)
			}
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
