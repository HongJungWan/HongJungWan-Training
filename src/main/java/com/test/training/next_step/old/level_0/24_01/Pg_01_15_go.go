package main

import "fmt"

func UpdateList(arr []int, flag []bool) []int {
	var X []int
	for i := 0; i < len(flag); i++ {
		if flag[i] {
			for j := 0; j < arr[i]*2; j++ {
				X = append(X, arr[i])
			}
		} else {
			X = X[:len(X)-arr[i]]
		}
	}
	return X
}

func main() {
	arr := []int{3, 2, 4, 1, 3}
	flag := []bool{true, false, true, false, false}
	result := UpdateList(arr, flag)
	fmt.Println(result)
}

// arr := []int{1, 2, 3, 4, 5}

// slice1 := arr[:3]  [1, 2, 3]
// slice2 := arr[1:4]  [2, 3, 4]
// slice3 := arr[2:]  [3, 4, 5]
