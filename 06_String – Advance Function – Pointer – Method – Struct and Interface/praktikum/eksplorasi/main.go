package main

import (
	"fmt"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	var offset int = 5
	// var lowercase string = strings.ToLower(s.name)

	for _, val := range s.name {
		if val >= 'a' && val <= 'z' {
			char := 'a' + (val-'a'+rune(offset))%26
			nameEncode += string(char)
		} else if val >= 'A' && val <= 'Z' {
			char := 'A' + (val-'A'+rune(offset))%26
			nameEncode += string(char)
		} else {
			nameEncode += string(val)
		}
	}

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	var offset int = 5
	// var lowercase string = strings.ToLower(s.name)

	for _, val := range s.name {
		if val >= 'a' && val <= 'z' {
			char := 'a' + (val-'a'-rune(offset))%26
			nameDecode += string(char)
		} else if val >= 'A' && val <= 'Z' {
			char := 'A' + (val-'A'-rune(offset))%26
			nameDecode += string(char)
		} else {
			nameDecode += string(val)
		}
	}

	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student's name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student's name " + a.name + " is : " + c.Decode())
	}
}
