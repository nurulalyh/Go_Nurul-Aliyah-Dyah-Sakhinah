package main

import (
	"fmt"
)

func Compare(a, b string) string {
	var compareResult string

	var minimumLen int = len(a)
	if len(b) < minimumLen {
		minimumLen = len(b)
	}

	for i := 0; i < minimumLen; i++ {
		if a[i] == b[i] {
			compareResult += string(a[i])
		} else {
			break
		}
	}

	return compareResult
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     //AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  //KANG
	fmt.Println(Compare("KI", "KIJANG"))      //KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    //ILA
}