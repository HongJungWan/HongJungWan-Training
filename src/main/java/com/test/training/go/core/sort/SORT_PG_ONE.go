// PG - K번째수
package main

import (
	"fmt"
	"sort"
)

func findKthElements(array []int, commands [][]int) []int {
	var answer []int

	for _, command := range commands {
		i, j, k := command[0], command[1], command[2]

		slice := make([]int, j-i+1)
		copy(slice, array[i-1:j])
		sort.Ints(slice)

		answer = append(answer, slice[k-1])
	}
	return answer
}

func main() {
	array := []int{1, 5, 2, 6, 3, 7, 4}
	commands := [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}

	result := findKthElements(array, commands)
	fmt.Println(result) // [5, 6, 3] 출력
}

// 로직

// 1. 배열을 i부터 j까지 자른 후 새 슬라이스에 저장
// 2. 슬라이스 정렬
// 3. k번째 숫자 찾아서 결과에 추가
