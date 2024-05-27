// TODO: 아나그램 체크 문제
// TODO: 두 단어가 같은 문자를 같은 개수만큼 가지고 있어, 그 문자의 순서를 바꾸면 서로 같은 단어가 될 수 있는 관계

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	word1 := scanner.Text()

	scanner.Scan()
	word2 := scanner.Text()

	fmt.Println(minDeletionsToMakeAnagram(word1, word2))
}

func minDeletionsToMakeAnagram(word1, word2 string) int {
	// 두 단어의 문자 빈도수를 계산할 배열
	count1 := make([]int, 26)
	count2 := make([]int, 26)

	// 첫 번째 단어의 문자 빈도수 계산
	for _, ch := range word1 {
		count1[ch-'a']++
	}

	// 두 번째 단어의 문자 빈도수 계산
	for _, ch := range word2 {
		count2[ch-'a']++
	}

	// 삭제해야 하는 문자 개수 계산
	deletions := 0
	for i := 0; i < 26; i++ {
		deletions += abs(count1[i] - count2[i])
	}
	return deletions
}

// 절대값 계산 함수
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
