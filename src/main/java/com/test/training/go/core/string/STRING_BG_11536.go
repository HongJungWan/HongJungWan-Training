// 줄 세우기

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text()) // 정수로 변환

	names := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		names[i] = scanner.Text()
	}

	if isIncreasing(names) {
		fmt.Println("INCREASING")
	} else if isDecreasing(names) {
		fmt.Println("DECREASING")
	} else {
		fmt.Println("NEITHER")
	}
}

func isIncreasing(arr []string) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			return false
		}
	}
	return true
}

func isDecreasing(arr []string) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] <= arr[i] {
			return false
		}
	}
	return true
}
