package main

import (
	"fmt"
)

func Pg_11_20(arr []int) []int {
	var X []int
	for _, a := range arr {
		for i := 0; i < a; i++ {
			X = append(X, a)
		}
	}
	return X
}

func main() {
	fmt.Println(Pg_11_20([]int{5, 1, 4})) // [5, 5, 5, 5, 5, 1, 4, 4, 4, 4]
	fmt.Println(Pg_11_20([]int{6, 6}))    // [6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6]
	fmt.Println(Pg_11_20([]int{1}))       // [1]
}
