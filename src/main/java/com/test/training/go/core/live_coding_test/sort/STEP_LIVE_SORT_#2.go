// TODO: 삽입 정렬 (Insertion Sort)
// TODO: 앞에서부터 차례대로 이미 정렬된 배열 부분과 비교하여, 자신의 위치를 찾은 후 삽입하여 정렬

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	insertionSortItems := strings.Split(input, " ")
	numbers := make([]int, len(insertionSortItems))

	for i, s := range insertionSortItems {
		number, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}
		numbers[i] = number
	}

	insertionSort(numbers)
	fmt.Println("Sorted numbers:", numbers)
}

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

// Test Case : 4 5 1
