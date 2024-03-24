// PG - 의상
// 각 의상 종류별로 가능한 조합의 수를 계산하는 문제
package main

import "fmt"

func countClothingCombinations(clothes [][]string) int {
	clothesMap := make(map[string]int)

	for _, item := range clothes {
		clothesMap[item[1]]++
	}

	combinations := 1
	for _, count := range clothesMap {
		combinations *= (count + 1)
	}
	return combinations - 1
}

func main() {
	// 테스트 데이터
	clothes1 := [][]string{{"yellow_hat", "headgear"}, {"blue_sunglasses", "eyewear"}, {"green_turban", "headgear"}}

	// 결과 출력
	fmt.Println(countClothingCombinations(clothes1)) // 5
}

// 로직
// 1. 의상 종류별로 카운트를 저장할 해시맵
// 2. 해당 종류의 의상 수를 카운트
// 3. 모든 조합의 수를 계산 (각 종류별 의상 수 + 안 입는 경우)

// 조건 1 : 최소 한 개의 의상은 입으므로, 모든 종류를 안 입는 경우를 제외
