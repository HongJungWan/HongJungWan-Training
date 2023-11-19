package main

import (
	"fmt"
)

func Pg_11_14(myString string) string {
	result := ""
	for _, char := range myString {
		if char < 'l' {
			result += "l"
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	test1 := Pg_11_14("abcdevwxyz")
	test2 := Pg_11_14("jjnnllkkmm")

	fmt.Println(test1) // Output: "lllllvwxyz"
	fmt.Println(test2) // Output: "llnnllllmm"
}
