// TODO: 문자열 뒤집기

package main

import (
	"fmt"
)

func main() {
	input := "hello"
	reversed := reverseString(input)

	fmt.Println("Original:", input)
	fmt.Println("Reversed:", reversed)
}

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	// 문자열을 반으로 나누어 양쪽 끝에서부터 바꾸기
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}
