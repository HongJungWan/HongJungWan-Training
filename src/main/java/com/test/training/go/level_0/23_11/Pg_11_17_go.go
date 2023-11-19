package main

import (
	"fmt"
	"strings"
)

// 문자열에서 'A'와 'B'를 서로 바꾸는 함수
func swapAB(str string) string {
	var swapped strings.Builder
	for _, char := range str {
		switch char {
		case 'A':
			swapped.WriteRune('B')
		case 'B':
			swapped.WriteRune('A')
		default:
			swapped.WriteRune(char)
		}
	}
	return swapped.String()
}

func Pg_11_17(myString, pat string) int {
	swappedString := swapAB(myString)
	if strings.Contains(swappedString, pat) {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(Pg_11_17("ABBAA", "AABB")) // 1
	fmt.Println(Pg_11_17("ABAB", "ABAB"))  // 0
}
