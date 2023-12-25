// PG - 가장 큰 수
package main

import (
	"fmt"
	"sort"
	"strconv"
)

type ByLargestNumber []string

func (a ByLargestNumber) Len() int { return len(a) }

func (a ByLargestNumber) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByLargestNumber) Less(i, j int) bool {
	return a[i]+a[j] > a[j]+a[i]
}

func largestNumber(numbers []int) string {
	strNumbers := make([]string, len(numbers))
	for i, number := range numbers {
		strNumbers[i] = strconv.Itoa(number)
	}

	sort.Sort(ByLargestNumber(strNumbers))

	if strNumbers[0] == "0" {
		return "0"
	}

	result := ""
	for _, number := range strNumbers {
		result += number
	}

	return result
}

func main() {
	// 테스트
	fmt.Println(largestNumber([]int{6, 10, 2}))        // "6210"
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9})) // "9534330"
}
