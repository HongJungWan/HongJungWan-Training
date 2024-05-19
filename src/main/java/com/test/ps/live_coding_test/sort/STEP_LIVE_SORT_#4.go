// TODO: 병합 정렬 (Merge Sort)
// TODO: 배열을 반씩 나누어 재귀적으로 정렬한 후, 두 정렬된 부분 배열을 합쳐 전체 배열을 정렬

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

	mergeSortItems := strings.Split(input, " ")
	numbers := make([]int, len(mergeSortItems))

	for i, s := range mergeSortItems {
		number, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}
		numbers[i] = number
	}

	mergeSort(numbers, 0, len(numbers)-1)
	fmt.Println("Sorted numbers:", numbers)
}

func mergeSort(arr []int, left int, right int) {
	if left < right {
		mid := left + (right-left)/2

		// 좌측과 우측을 재귀적으로 정렬
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)

		// 병합 실행
		merge(arr, left, mid, right)
	}
}

// 두 부분 배열을 병합하는 함수
func merge(arr []int, left int, mid int, right int) {
	n1 := mid - left + 1
	n2 := right - mid

	// 임시 배열 생성
	L := make([]int, n1)
	R := make([]int, n2)

	// 데이터를 임시 배열에 복사
	for i := 0; i < n1; i++ {
		L[i] = arr[left+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[mid+1+j]
	}

	// 임시 배열을 병합하여 원래 배열에 복사
	i, j, k := 0, 0, left
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	// 잔여 요소들을 원래 배열에 복사
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}
