package main

import (
	"fmt"
	"strings"
)

func Pg_12_10(strArr []string) []string {
	var result []string
	for _, str := range strArr {
		if !strings.Contains(str, "ad") {
			result = append(result, str)
		}
	}
	return result
}

func main() {
	example1 := []string{"and", "notad", "abcd"}
	example2 := []string{"there", "are", "no", "a", "ds"}

	result1 := Pg_12_10(example1)
	result2 := Pg_12_10(example2)

	fmt.Println(result1) // Output: ["and", "abcd"]
	fmt.Println(result2) // Output: ["there", "are", "no", "a", "ds"]
}
