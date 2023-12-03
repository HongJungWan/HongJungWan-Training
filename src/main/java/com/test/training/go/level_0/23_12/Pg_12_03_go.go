package main

import (
	"fmt"
)

func Pg_12_03(numList []int) int {
	evenSum := 0
	oddSum := 0

	for i, num := range numList {
		if i%2 == 0 {
			oddSum += num
		} else {
			evenSum += num
		}
	}
	if evenSum > oddSum {
		return evenSum
	}
	return oddSum
}

func main() {
	numListOne := []int{4, 2, 6, 1, 7, 6}
	fmt.Println(Pg_12_03(numListOne))

	numListTwo := []int{-1, 2, 5, 6, 3}
	fmt.Println(Pg_12_03(numListTwo))
}
