package main

import (
	"fmt"
)

func Pg_11_16(arr []int, deleteList []int) []int {
	deleteMap := make(map[int]bool)
	for _, item := range deleteList {
		deleteMap[item] = true
	}

	var result []int
	for _, item := range arr {
		if !deleteMap[item] {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	arr1 := []int{293, 1000, 395, 678, 94}
	deleteList1 := []int{94, 777, 104, 1000, 1, 12}
	fmt.Println("Result 1:", Pg_11_16(arr1, deleteList1))

	arr2 := []int{110, 66, 439, 785, 1}
	deleteList2 := []int{377, 823, 119, 43}
	fmt.Println("Result 2:", Pg_11_16(arr2, deleteList2))
}
