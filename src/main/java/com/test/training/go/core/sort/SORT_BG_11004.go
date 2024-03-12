package main

import (
	"fmt"
)

// Quick Select 알고리즘 구현
func quickSelect(arr []int, left, right, K int) int {
	if left <= right {
		partitionIndex := partition(arr, left, right)

		if partitionIndex == K {
			return arr[partitionIndex]
		} else if partitionIndex < K {
			return quickSelect(arr, partitionIndex+1, right, K)
		} else {
			return quickSelect(arr, left, partitionIndex-1, K)
		}
	}
	return -1 // 이 경우는 발생하지 않음
}

// Quick Sort의 파티션 함수
func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	nums := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	// K-1을 전달하는 이유는 배열의 인덱스가 0부터 시작하기 때문.
	result := quickSelect(nums, 0, N-1, K-1)
	fmt.Println(result)
}

// 취미용 - 깊게 보진 말고 살짝~
