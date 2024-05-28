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
	wordOne := scanner.Text()

	scanner.Scan()
	wordTwo := scanner.Text()

	fmt.Println(minDeletionsToMakeAnagram(wordOne, wordTwo))
}

func minDeletionsToMakeAnagram(wordOne, wordTwo string) int {
	// 두 단어의 문자 빈도수를 계산할 배열
	count1 := make([]int, 26)
	count2 := make([]int, 26)

	// 첫 번째 단어의 문자 빈도수 계산
	for _, ch := range wordOne {
		count1[ch-'a']++
	}

	// 두 번째 단어의 문자 빈도수 계산
	for _, ch := range wordTwo {
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

/*
두 개의 영어 단어가 주어졌을 때, 두 단어가 서로 애너그램 관계에 있도록 만들기 위해서 제거해야 하는 최소 개수의 문자 수를 구하는 프로그램을 작성하시오.
문자를 제거할 때에는 아무 위치에 있는 문자든지 제거할 수 있다.

입력
aabbcc
xxyybb

출력
8

---
wordOne = "aabbcc"
count1 = [2, 2, 2, 0, 0, 0, ..., 0]

wordTwo = "xxyybb"
count2 = [0, 2, 0, 0, 0, 0, ..., 2, 2]


두 배열의 차이 계산 및 삭제 개수 합산:
'a': 2 - 0 = 2
'b': 2 - 2 = 0
'c': 2 - 0 = 2
'x': 0 - 2 = 2
'y': 0 - 2 = 2
나머지 알파벳: 차이 없음 = 0

최종적으로 deletions 값은 2 + 0 + 2 + 2 + 2 = 8
*/
