// BG - 동일한 단어 그룹화하기

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 문자열 정렬
func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	N := 0
	fmt.Sscanf(scanner.Text(), "%d", &N)

	wordMap := make(map[string]bool)

	for i := 0; i < N; i++ {
		scanner.Scan()
		word := scanner.Text()
		sortedWord := sortString(word)
		wordMap[sortedWord] = true
	}

	fmt.Println(len(wordMap))
}

// sort.Slice는 제공된 슬라이스를 정렬하는 함수.
// 여기서 r은 정렬할 룬 슬라이스.
// 이 함수는 두 번째 매개변수로 "less" 함수를 받는데, 이 함수는 슬라이스의 두 요소 i와 j가 주어졌을 때,
// r[i]가 r[j]보다 "작은지" (즉, 알파벳 순서에서 먼저 오는지)를 결정.
// 이 비교 함수를 기반으로 sort.Slice는 슬라이스를 오름차순으로 정렬
