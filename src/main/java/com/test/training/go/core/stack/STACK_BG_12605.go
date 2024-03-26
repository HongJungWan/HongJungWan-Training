package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 입력을 받기 위한 bufio.Scanner 생성
	scanner := bufio.NewScanner(os.Stdin)

	// 첫 번째 줄 (케이스의 개수 N) 읽기
	scanner.Scan()
	var N int
	fmt.Sscanf(scanner.Text(), "%d", &N)

	// N개의 케이스를 처리
	for i := 1; i <= N; i++ {
		scanner.Scan()
		// 현재 케이스의 단어들을 공백으로 분리
		words := strings.Split(scanner.Text(), " ")
		// 단어들을 반대 순서로 뒤집기
		for j, k := 0, len(words)-1; j < k; j, k = j+1, k-1 {
			words[j], words[k] = words[k], words[j]
		}
		// 결과 출력
		fmt.Printf("Case #%d: %s\n", i, strings.Join(words, " "))
	}
}
