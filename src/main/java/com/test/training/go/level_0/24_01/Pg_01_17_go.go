package main

import (
	"fmt"
)

func RemoveCharsByIndices(myString string, indices []int) string {
	result := ""

	for i := 0; i < len(myString); i++ {
		if Contains(indices, i) {
			result += string(myString[i])
		}
	}
	return result
}

func Contains(arr []int, num int) bool {
	for _, value := range arr {
		if value == num {
			return false
		}
	}
	return true
}

func main() {
	myString := "apporoograpemmemprs"
	indices := []int{1, 16, 6, 15, 0, 10, 11, 3}
	fmt.Println(RemoveCharsByIndices(myString, indices)) // "programmers"
}
