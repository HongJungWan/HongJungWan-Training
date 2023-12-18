package main

import (
	"fmt"
)

func Pg_12_18(str1, str2 string) string {
	var result string
	for index := range str1 {
		result += string(str1[index]) + string(str2[index])
	}
	return result
}

func main() {
	str1 := "aaaaa"
	str2 := "bbbbb"
	fmt.Println(Pg_12_18(str1, str2)) // "ababababab"
}
