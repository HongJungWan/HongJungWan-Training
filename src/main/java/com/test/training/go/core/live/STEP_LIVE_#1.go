// TODO: 선택 정렬 (Selection Sort)
// TODO: 배열에서 최소값을 찾아 선택한 다음, 배열의 앞부분으로 이동시켜서 정렬

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

	selectionSortItems := strings.Split(input, " ")
	numbers := make([]int, len(selectionSortItems))

	for i, s := range selectionSortItems {
		number, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}
		numbers[i] = number
	}

	selectionSort(numbers)
	fmt.Println("Sorted numbers:", numbers)
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// Test Case : 4 5 1
