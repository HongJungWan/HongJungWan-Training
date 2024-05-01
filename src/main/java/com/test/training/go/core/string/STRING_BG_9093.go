package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T := scanner.Text()

	// 문자열로부터 정수 변환
	testCaseCount, err := strconv.Atoi(T)
	if err != nil {
		fmt.Println("Error reading the number of test cases:", err)
		return
	}

	for i := 0; i < testCaseCount; i++ {
		scanner.Scan()
		sentence := scanner.Text()
		words := strings.Split(sentence, " ")
		for j, word := range words {
			// 단어 뒤집기
			words[j] = reverse(word)
		}
		// 결과 출력
		fmt.Println(strings.Join(words, " "))
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
