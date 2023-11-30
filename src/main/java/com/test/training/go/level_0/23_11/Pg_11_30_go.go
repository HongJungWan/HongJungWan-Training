package main

import "fmt"

func solution(num_list []int, n int) []int {
	firstPart := num_list[n:]
	secondPart := num_list[:n]

	return append(firstPart, secondPart...)
}

func main() {
	num_list1 := []int{2, 1, 6}
	n1 := 1
	result1 := solution(num_list1, n1)
	fmt.Println(result1) // [1, 6, 2]

	num_list2 := []int{5, 2, 1, 7, 5}
	n2 := 3
	result2 := solution(num_list2, n2)
	fmt.Println(result2) // [7, 5, 5, 2, 1]
}
