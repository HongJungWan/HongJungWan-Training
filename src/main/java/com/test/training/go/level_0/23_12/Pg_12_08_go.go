package main

import "fmt"

func solution(arr []int, idx int) int {
	for i := idx + 1; i < len(arr); i++ {
		if arr[i] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	arr1 := []int{0, 0, 0, 1}
	idx1 := 1
	result1 := solution(arr1, idx1)
	fmt.Println(result1) // 출력: 3

	arr2 := []int{1, 0, 0, 1, 0, 0}
	idx2 := 4
	result2 := solution(arr2, idx2)
	fmt.Println(result2) // 출력: -1
}
