package main

import (
	"fmt"
)

func Pg_12_12(arr []int, queries [][]int) []int {
	n := len(arr)
	window := make([]int, n)

	for _, query := range queries {
		s, e := query[0], query[1]
		window[s]++
		if e+1 < n {
			window[e+1]--
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += window[i]
		arr[i] += sum
	}
	return arr
}

func main() {
	arr := []int{0, 1, 2, 3, 4}
	queries := [][]int{{0, 1}, {1, 2}, {2, 3}}

	fmt.Println(Pg_12_12(arr, queries)) // Output: [1, 3, 4, 4, 4]
}

// Sliding Window
