package main

import "fmt"

func main() {
	fmt.Println("\t\t\t\t\t\t\t\t\tProgram Fizz-Buzz")
	fmt.Println("")
    for i := 1; i <= 100; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Print("FizzBuzz ")
        } else if i%3 == 0 {
            fmt.Print("Fizz ")
        } else if i%5 == 0 {
            fmt.Print("Buzz ")
        } else {
            fmt.Printf("%d ", i)
        }
    }
}