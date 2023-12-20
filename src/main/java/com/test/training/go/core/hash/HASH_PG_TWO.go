// PG - 폰켓몬
package main

import "fmt"

func solution(nums []int) int {
	max := len(nums) / 2
	set := make(map[int]bool)

	for _, num := range nums {
		set[num] = true
	}

	if len(set) > max {
		return max
	} else {
		return len(set)
	}
}

func main() {
	fmt.Println(solution([]int{3, 1, 2, 3}))       // 출력: 2
	fmt.Println(solution([]int{3, 3, 3, 2, 2, 4})) // 출력: 3
	fmt.Println(solution([]int{3, 3, 3, 2, 2, 2})) // 출력: 2
}

// 로직

// 1. 최대 담을 수 있는 폰켓몬 수 계산

// 2. Set 역할을 하는 Map 생성 및 초기화
//
// Go 언어에는 집합(Set) 자료구조가 없으므로, map을 사용하여 유사한 기능을 구현.
// set이라는 이름의 map[int]bool 타입을 생성.
// 각 폰켓몬 종류의 번호를 키(Key)로, 그 값(Value)을 true로 설정.

// 3. 최대 선택 가능한 폰켓몬 수와 종류 수 비교
//
// map에 저장된 폰켓몬 종류의 수(len(set))와 최대로 선택할 수 있는 폰켓몬의 수(max)를 비교.
