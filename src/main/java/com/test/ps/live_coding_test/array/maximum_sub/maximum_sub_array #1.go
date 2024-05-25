// TODO: 최대 부분 합
// TODO: Kadane's Algorithm, 최대 부분 배열의 합을 찾는 'maxSubArray' 함수 사용

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
	T, _ := strconv.Atoi(scanner.Text())

	for t := 0; t < T; t++ {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		numsStr := strings.Split(scanner.Text(), " ")

		nums := make([]int, N)
		for i := 0; i < N; i++ {
			nums[i], _ = strconv.Atoi(numsStr[i])
		}
		fmt.Println(maxSubArray(nums))
	}
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

/*
입력
2 (2개의 테스트 케이스)
5 (배열의 크기)
1 2 3 4 5 (element)
5 (배열의 크기)
2 1 -2 3 -5 (element)

출력
15
4

아이디어
1. 각 테스트 케이스에 대해 배열 크기 N과 배열 요소를 읽는다.
2. Kadane's Algorithm, 최대 부분 배열의 합을 찾는 'maxSubArray' 함수 사용
3. 각 테스트 케이스에 대해 최대 부분 배열의 합을 출력
*/
