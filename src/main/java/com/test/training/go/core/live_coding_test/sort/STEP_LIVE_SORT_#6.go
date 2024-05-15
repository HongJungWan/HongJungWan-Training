// TODO: 퀵 정렬 (Quick Sort)
// TODO: 피벗을 기준으로 배열을 분할하여 피벗보다 작은 요소와 큰 요소로 나눈 뒤, 이 과정을 재귀적으로 반복해 정렬

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

	quickSortItems := strings.Split(input, " ")
	numbers := make([]int, len(quickSortItems))

	for i, s := range quickSortItems {
		number, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}
		numbers[i] = number
	}

	quickSort(numbers, 0, len(numbers)-1)
	fmt.Println("Sorted numbers:", numbers)
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		// pi는 파티셔닝 인덱스, arr[pi]는 이제 정확한 위치에 있다.
		pi := partition(arr, low, high)

		// 별도로 pi 전후의 요소들을 재귀적으로 정렬
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high] // 피벗
	i := low - 1       // 가장 작은 요소의 인덱스

	for j := low; j < high; j++ {
		// 현재 요소가 피벗보다 작거나 같은 경우
		if arr[j] <= pivot {
			i++ // 작은 요소의 인덱스 증가
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
