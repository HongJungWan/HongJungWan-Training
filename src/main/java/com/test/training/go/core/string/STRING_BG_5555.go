// BG - 반지

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 주어진 반지들 중에서 찾고자하는 문자열을 포함하는 반지의 수 계산
func countRingsContainingTarget(target string, rings []string) int {
	count := 0
	for _, ring := range rings {
		ring += ring // 순환 구조 생성
		if strings.Contains(ring, target) {
			count++
		}
	}
	return count
}

func readInput() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)

	// 찾고자 하는 문자열 읽기
	scanner.Scan()
	target := scanner.Text()

	// 반지의 개수 읽기
	scanner.Scan()
	var n int
	fmt.Sscanf(scanner.Text(), "%d", &n)

	// 반지 문자열들 읽기
	rings := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		rings[i] = scanner.Text()
	}

	return target, rings
}

func main() {
	target, rings := readInput()
	count := countRingsContainingTarget(target, rings)
	fmt.Println(count)
}

// [Go 언어 표준 라이브러리]

// strings.Contains(s, substr string) bool 형태로 선언
//
// 이 함수는 s라는 문자열 내에 substr이라는 하위 문자열이 포함되어 있는지를 검사
// 만약 substr이 s 안에 있으면 true를, 그렇지 않으면 false를 반환
