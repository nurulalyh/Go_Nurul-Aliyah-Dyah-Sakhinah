package main

import "fmt"

func main() {
    var a, b, t, L float32
    var tryAgain string

    for {
        fmt.Println("\t\tProgram Menghitung Luas Trapesium")
        fmt.Println("------------------------------------------------------------------")
        fmt.Print("Masukkan panjang sisi a (cm): ")
        fmt.Scanln(&a)
        fmt.Println("")
        fmt.Print("Masukkan panjang sisi b (cm): ")
        fmt.Scanln(&b)
        fmt.Println("")
        fmt.Print("Masukkan tinggi trapesium (cm): ")
        fmt.Scanln(&t)
        fmt.Println("------------------------------------------------------------------")

        L = 0.5 * (a + b) * t
        fmt.Printf("Luas Trapesium adalah %v cm\n", L)
		fmt.Println("------------------------------------------------------------------")

        fmt.Print("Apakah Anda ingin menghitung luas trapesium lagi? (y/n): ")
        fmt.Scanln(&tryAgain)
		fmt.Println("------------------------------------------------------------------")
        if tryAgain != "y" && tryAgain != "Y" {
            fmt.Println("Terima Kasih Telah Menggunakan Program Ini 😊")
			fmt.Println("------------------------------------------------------------------")
			break
        }
    }
}