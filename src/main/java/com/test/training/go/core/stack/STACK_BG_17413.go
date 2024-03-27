package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // 입력 문자열의 앞뒤 공백 제거

	fmt.Println(processInput(input))
}

// processInput 함수는 문자열을 처리하고 결과를 반환
func processInput(input string) string {
	var result strings.Builder // 결과를 저장할 Builder 생성
	var word strings.Builder   // 단어를 임시로 저장할 Builder 생성
	tag := false               // 태그 안에 있는지 확인하는 변수

	for _, ch := range input {
		if ch == '<' {
			result.WriteString(reverse(word.String()))
			word.Reset()
			tag = true
			result.WriteRune(ch)
		} else if ch == '>' {
			tag = false
			result.WriteRune(ch)
		} else if tag {
			result.WriteRune(ch)
		} else {
			if ch == ' ' {
				result.WriteString(reverse(word.String()))
				word.Reset()
				result.WriteRune(ch)
			} else {
				word.WriteRune(ch)
			}
		}
	}
	result.WriteString(reverse(word.String())) // 마지막 단어를 처리
	return result.String()
}

// reverse 함수는 주어진 문자열을 뒤집어서 반환
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
