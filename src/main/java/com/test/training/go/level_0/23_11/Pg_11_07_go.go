package main

import "fmt"

func Pg_11_07(arr []int, delete_list []int) []int {
	delete_map := make(map[int]bool)
	for _, val := range delete_list {
		delete_map[val] = true
	}

	var result []int

	for _, val := range arr {
		if !delete_map[val] {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	//given
	arr1 := []int{293, 1000, 395, 678, 94}
	delete_list1 := []int{94, 777, 104, 1000, 1, 12}

	arr2 := []int{110, 66, 439, 785, 1}
	delete_list2 := []int{377, 823, 119, 43}

	//when
	result1 := Pg_11_07(arr1, delete_list1)
	result2 := Pg_11_07(arr2, delete_list2)

	//then
	fmt.Println()
	fmt.Println(result1)
	fmt.Println(result2)
}
