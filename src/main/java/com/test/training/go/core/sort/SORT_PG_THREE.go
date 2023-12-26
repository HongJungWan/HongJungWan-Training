// PG - H-Index
package main

import (
	"fmt"
	"sort"
)

func calculate(citations []int) int {
	sort.Ints(citations)
	n := len(citations)

	hIndex := 0
	for i := n - 1; i >= 0; i-- {
		if citations[i] > i {
			hIndex++
		} else {
			break
		}
	}
	return hIndex
}

func main() {
	citations := []int{3, 0, 6, 1, 5}
	fmt.Println(calculate(citations)) // 3
}
