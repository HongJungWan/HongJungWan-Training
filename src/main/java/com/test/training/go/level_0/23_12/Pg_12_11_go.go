package main

import "fmt"

func Pg_12_11(n int) [][]int {
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
		arr[i][i] = 1
	}
	return arr
}

func main() {
	testCases := []int{3, 6, 1}
	for _, n := range testCases {
		result := Pg_12_11(n)
		fmt.Println("n =", n, "Result:", result)
	}
}
