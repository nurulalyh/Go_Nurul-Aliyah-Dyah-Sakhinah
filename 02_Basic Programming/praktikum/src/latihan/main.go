package main

import "fmt"

func main(){
	// VARIABEL DAN TIPE DATA
	// var usia int
	// var nickname string = "lulu"
	// var firstname, lastname = "Nurul", "Aliyah"
	// gender := "wanita"
	// usia = 16
	// const pi float64 = 3.14
	// const (
	// 	firstname string = "Nurul"
	// 	lastname string = "Aliyah"
	// 	usia int = 16
	// 	gender = "wanita"
	// )
	// var islogin bool = false

	// fmt.Println("Full Name :", firstname, lastname)
	// fmt.Println(lastname)
	// fmt.Println(nickname)
	// fmt.Println(gender)
	// fmt.Println(usia)
	// fmt.Println(islogin)

	// //OPERAND DAN OPERATOR
	// x := 1 + 2 //operand adalah angka, operator adalah +
	// fmt.Println(x)

	// //EXPRESSION
	// a := 5
	// b := 6
	// c := a + b //c adalah expression, digunakan agar program dapat berjalan lebih baik

	// fmt.Println(c)

	// // LUAS SEGITIGA

	// a := 10
	// t := 15
	// L := (a * t) / 2
	// fmt.Println("Luas Segitiga Adalah", L)
	
	// // STRING OPERATION
	// helloworld := "Hello" + " " + "World"
	// fmt.Println(helloworld)

	// // IF, ELSE IF, ELSE
	// hour := 20

	// if hour >= 3 && hour < 12 {
	// 	fmt.Println("Selamat Pagi")
	// } else if hour >= 12 && hour < 18 {
	// 	fmt.Println("Selamat Sore")
	// } else {
	// 	fmt.Println("Selamat Malam")
	// }

	// // SWITCH
	// var today int = 7

	// switch today {
	// case 1:
	// 	fmt.Println("Hari ini Senin")
	// case 2:
	// 	fmt.Println("Hari ini Selasa")
	// case 3:
	// 	fmt.Println("Hari ini Rabu")
	// case 4:
	// 	fmt.Println("Hari ini Kamis")
	// case 5:
	// 	fmt.Println("Hari ini Jum'at")
	// case 6:
	// 	fmt.Println("Hari ini Sabtu")
	// case 7:
	// 	fmt.Println("Hari ini Minggu")
	// default:
	// 	fmt.Println("Invalid date")
	// }

	// // LOOPING for
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// // LOOPING Continue (memaksakan perulangan berlanjut atau diskip ke perulangan selanjurnya) & Break (kebalikan continue)
	// for i := 0; i < 5; i++ {
	// 	if i == 1 {
	// 		continue
	// 	}
	// 	if i > 3 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	//LOOPING OVER STRING
	sentence := "Hello"

	for i := 0; i < len(sentence); i++ {
		fmt.Printf(string(sentence[i]) + "-") //Println otomatis enter, Printf lanjut kesebelahnya
	}
	fmt.Println("      ----->>>")
	for pos, char := range sentence{
		fmt.Printf("character %c start at byte position %d\n", char, pos) 
		// %c Digunakan untuk memformat data numerik yang merupakan kode unicode, menjadi bentuk string karakter unicode-nya.
		// %d Digunakan untuk memformat data numerik, menjadi bentuk string numerik berbasis 10
	}
}