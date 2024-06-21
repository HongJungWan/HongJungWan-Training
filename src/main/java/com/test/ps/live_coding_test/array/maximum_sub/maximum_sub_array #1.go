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

핵심
1. 전체 배열이 주어졌고, 우리는 부분 배열이 가질 수 있는 부분합들 중에서 최대 부분합을 구해야 한다.
2. O(N)으로 풀기 위한 핵심은 [각각의 최대 부분합]은 [이전 최대 부분합이 반영된 결과값]이라는 것이다.
*/

/*
input : [2, 1, -2, 3, -5]

1. 초기화 : maxSum = 2, currentSum = 2 (첫 번째 요소)

2. 첫 번째 반복 (i = 1, nums[i] = 1)
2-1. currentSum이 양수이므로, currentSum += nums[i] = 2 + 1 = 3
2-2. maxSum 업데이트: maxSum = max(3, 2) = 3

3. 두 번째 반복 (i = 2, nums[i] = -2)
3-1. currentSum이 여전히 양수이므로, currentSum += nums[i] = 3 - 2 = 1
3-2. maxSum 업데이트: maxSum = max(3, 1) = 3

4. 세 번째 반복 (`i = 3, nums[i] = 3`)
4-1. currentSum이 양수이므로, currentSum += nums[i] = 1 + 3 = 4
4-2. maxSum 업데이트: maxSum = max(4, 3) = 4

5. 네 번째 반복 (i = 4, nums[i] = -5)
5-1. currentSum이 여전히 양수이므로, currentSum += nums[i] = 4 - 5 = -1
5-2. maxSum 업데이트: maxSum = max(4, -1) = 4
*/
