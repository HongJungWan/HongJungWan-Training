package main

import "fmt"

func pg_12_29(a int, d int, included []bool) int {
	count := a
	result := 0

	for i := 0; i < len(included); i++ {
		if included[i] == true {
			result += count
		}
		count += d
	}
	return result
}

func main() {
	testBooleanData := []bool{true, false, false, true, true}
	fmt.Println(pg_12_29(3, 4, testBooleanData))
}
