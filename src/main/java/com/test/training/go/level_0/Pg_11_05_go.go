package main

import (
	"fmt"
	"sort"
)

func solution(num_list []int) []int {
	sort.Ints(num_list)

	return num_list[:5]
}

func main() {
	num_list := []int{12, 4, 15, 46, 38, 1, 14}
	result := solution(num_list)

	fmt.Println(result)
}
