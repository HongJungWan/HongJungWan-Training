package main

import (
	"fmt"
	"strings"
)

func Pg_11_04(my_string string) []string {
	// strings.Fields 함수를 사용하여 공백으로 문자열을 나눈다.
	words := strings.Fields(my_string)
	return words
}

func main() {
	example1 := " i    love  you"
	example2 := "    programmers  "

	fmt.Println()
	fmt.Println(Pg_11_04(example1)) // ["i", "love", "you"]
	fmt.Println(Pg_11_04(example2)) // ["programmers"]
}
