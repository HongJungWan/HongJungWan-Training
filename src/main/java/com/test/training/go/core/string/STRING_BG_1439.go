// BG - 뒤집기

package main

import (
	"bufio"
	"fmt"
	"os"
)

func minFlips(s string) int {
	// 첫 번째 문자를 기준으로 설정
	currentChar := s[0]
	// 0과 1의 덩어리 개수를 각각 세기
	count0, count1 := 0, 0

	for i := 0; i < len(s); i++ {
		// 문자가 바뀌면 덩어리 개수 증가
		if s[i] != currentChar {
			if currentChar == '0' {
				count0++
			} else {
				count1++
			}
			currentChar = s[i]
		}
	}

	// 마지막 덩어리 처리
	if currentChar == '0' {
		count0++
	} else {
		count1++
	}

	// 더 작은 덩어리의 개수 반환
	if count0 < count1 {
		return count0
	}
	return count1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // 문자열 S를 입력받음
	s := scanner.Text()

	fmt.Println(minFlips(s))
}
