package main

import (
	"fmt"
)

func Pg_12_13(arr []int, queries [][]int) []int {
	n := len(arr)
	window := make([]int, n)

	for _, query := range queries {
		s, e := query[0], query[1]
		window[s]++

		// 차분 배열(Difference Array) 기법
		if e+1 < n {
			window[e+1]--
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += window[i]
		arr[i] += sum
	}
	return arr
}

func main() {
	arr := []int{0, 1, 2, 3, 4}
	queries := [][]int{{0, 1}, {1, 2}, {2, 3}}

	fmt.Println(Pg_12_13(arr, queries)) // Output: [1, 3, 4, 4, 4]
}

// "차분 배열 문제"

// 차분 배열 기법에서는, 범위 업데이트를 두 단계로 나누어 처리한다.
//
// 시작 지점에서 증가(Increment at Start): 범위의 시작 지점 s에서 값을 증가시켜, 해당 지점부터 끝점까지의 모든 요소에 영향을 미치게 한다.
// 종료 지점 다음에서 감소(Decrement After End): 범위의 종료 지점 e 다음 위치에서 값을 감소시켜, 범위 밖의 요소들에는 영향을 미치지 않게 한다.

// 위 로직의 중요성
// 경계 설정: if e+1 < n { window[e+1]-- }는 범위의 끝 지점 바로 다음에 감소를 적용한다. 범위를 초과하는 부분에 대한 업데이트가 적용되지 않도록 한다.
// 누적된 변경 제한: 이 로직 없이는 시작 지점 s에서 증가된 값이 배열의 끝까지 계속 적용된다.
