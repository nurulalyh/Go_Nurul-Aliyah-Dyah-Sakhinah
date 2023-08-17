package main

import "fmt"

func main(){
	var nilai float32
	var tryAgain string

	
    for {
		fmt.Println("\t\tProgram Menentukan Grade Nilai")
        fmt.Println("-------------------------------------------------------------")
        fmt.Print("Masukkan Nilai (ex : 90.00): ")
        fmt.Scanln(&nilai)
        fmt.Println("")
        fmt.Println("-------------------------------------------------------------")

        if nilai <= 100 && nilai >= 80 {
			fmt.Printf("Nilai anda %v, Grade anda adalah A\n", nilai)
			fmt.Println("-------------------------------------------------------------")
		} else if nilai < 80 && nilai >= 65 {
			fmt.Printf("Nilai anda %v, Grade anda adalah B\n", nilai)
			fmt.Println("-------------------------------------------------------------")
		} else if nilai < 65 && nilai >= 50 {
			fmt.Printf("Nilai anda %v, Grade anda adalah C\n", nilai)
			fmt.Println("-------------------------------------------------------------")
		} else if nilai < 50 && nilai >= 35 {
			fmt.Printf("Nilai anda %v, Grade anda adalah D\n", nilai)
			fmt.Println("-------------------------------------------------------------")
		} else if nilai < 35 && nilai >= 0 {
			fmt.Printf("Nilai anda %v, Grade anda adalah E\n", nilai)
			fmt.Println("-------------------------------------------------------------")
		} else {
			fmt.Println("Nilai Invalid")
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