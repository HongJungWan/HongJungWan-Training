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
		// 분할 (Divide)
		partitionIndex := partition(arr, low, high)

		// 정복 (Conquer)
		quickSort(arr, low, partitionIndex-1)
		quickSort(arr, partitionIndex+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	mid := low + (high-low)/2                 // 중간 인덱스 계산
	arr[mid], arr[high] = arr[high], arr[mid] // 중간 값을 끝으로 이동

	pivot := arr[high] // 이제 중간 값이 피벗
	i := low - 1       // 가장 작은 요소의 인덱스

	for j := low; j < high; j++ {
		// 현재 요소가 피벗보다 작거나 같은 경우
		if arr[j] <= pivot {
			i++ // 작은 요소의 인덱스 증가
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // 피벗을 올바른 위치로 교환
	return i + 1
}
