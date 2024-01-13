package main

import (
	"fmt"
	"strings"
)

func FindLongestSubstring(myString string, pat string) string {
	index := strings.LastIndex(myString, pat)
	if index == -1 {
		return ""
	}
	return myString[:index+len(pat)]
}

func main() {
	fmt.Println(FindLongestSubstring("AbCdEFG", "dE")) // "AbCdE"
	fmt.Println(FindLongestSubstring("AAAAaaaa", "a")) // "AAAAaaaa"
}

// 예시: myString이 "AbCdEFG"이고 pat이 "dE"라면, LastIndex는 "dE" 문자열이 시작되는 위치인 3 반환
// 예시: myString[:index+len(pat)] -> myString[:5], 0 ~4 까지의 문자들 추출, 슬라이싱 문법
