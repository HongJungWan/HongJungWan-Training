package main

import "fmt"

func Pg_12_14(arr []string) string {
	result := ""

	for _, word := range arr {
		result += word
	}
	return result
}

func main() {
	arr := []string{"a", "b", "c"}
	fmt.Println(Pg_12_14(arr))
}
