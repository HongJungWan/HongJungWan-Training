// TODO: Two Sum 문제
// TODO: 주어진 배열에서 두 숫자의 합이 특정 목표값(target)과 같아지는 두 숫자의 인덱스를 찾는 문제

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result) // 출력: [0, 1]
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return []int{}
}

/*
[첫 번째 반복] (i = 0, num = 2)

1-1. target - num 계산: 9 - 2 = 7
1-2. numMap에서 7을 찾기 시도: 실패
1-3. numMap에 현재 숫자 추가: numMap[2] = 0
1-4. 현재 상태: numMap = map[int]int{2: 0}

---

[두 번째 반복] (i = 1, num = 7)

1-1. target - num 계산: 9 - 7 = 2
1-2. numMap에서 2을 찾기 시도: 성공 (j = 0)
1-3. 결과 반환: [0, 1]
1-4. twoSum 함수 종료, result 값은 [0, 1]
*/

/*
numMap[target-num], 맵 numMap에서 키가 target-num인 element를 조회.

if j, ok := numMap[target-num]; ok { ... }, Go 언어의 "comma ok" 패턴, 맵 조회 시 특정 키가 존재하는지 확인.
*/
