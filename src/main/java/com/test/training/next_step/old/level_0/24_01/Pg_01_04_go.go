package main

import (
	"fmt"
)

func reverseSubstring(my_string string, s int, e int) string {
	// 문자열을 룬 슬라이스로 변환
	runes := []rune(my_string)

	for i, j := s, e; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(reverseSubstring("Progra21Sremm3", 6, 12))  // "ProgrammerS123"
	fmt.Println(reverseSubstring("Stanley1yelnatS", 4, 10)) // "Stanley1yelnatS"
}
