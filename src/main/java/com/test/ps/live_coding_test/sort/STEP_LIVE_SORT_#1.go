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

/* 아이디어

초기 상태: 9, 6, 7, 3, 5

1회전: [최솟값 탐색 : 3], 첫번째 값 9와 최솟값 3을 교환, [3, 6, 7, 9, 5]
2회전: [최솟값 탐색 : 5], 두번째 값 6과 최솟값 5를 교환, [3, 5, 7, 9, 6]

*/
