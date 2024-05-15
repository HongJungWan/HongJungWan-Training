// TODO: 힙 정렬 (Heap Sort)
// TODO: 최대 힙을 만든 후, 루트 요소를 반복적으로 제거하고 배열의 끝으로 옮겨 최종적으로 정렬된 배열을 얻는 방식.

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

	heapSortItems := strings.Split(input, " ")
	numbers := make([]int, len(heapSortItems))

	for i, s := range heapSortItems {
		number, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}
		numbers[i] = number
	}

	heapSort(numbers)
	fmt.Println("Sorted numbers:", numbers)
}

func heapSort(arr []int) {
	n := len(arr)

	// 배열을 최대 힙으로 구성
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 힙에서 하나씩 요소를 제거하여 배열의 끝으로 이동
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // 루트(최대값)를 배열의 끝으로 이동
		heapify(arr, i, 0)              // 힙 속성 유지
	}
}

// 주어진 노드 i를 루트로 하는 서브 트리를 힙 속성에 맞게 조정
func heapify(arr []int, n int, i int) {
	largest := i     // 가장 큰 요소를 루트로 초기화
	left := 2*i + 1  // 왼쪽 자식 인덱스
	right := 2*i + 2 // 오른쪽 자식 인덱스

	// 왼쪽 자식이 현재 요소보다 크면, 왼쪽 자식의 인덱스를 largest에 저장
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 오른쪽 자식이 현재 요소보다 크면, 오른쪽 자식의 인덱스를 largest에 저장
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// largest가 기존의 i와 다르다면, 요소의 위치를 바꾸고 힙 속성을 재조정
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}
