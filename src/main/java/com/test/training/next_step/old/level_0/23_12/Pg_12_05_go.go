package main

import (
	"fmt"
	"strings"
)

func Pg_12_05(myString string) []int {
	splitStrings := strings.Split(myString, "x")

	var lengths []int
	for _, s := range splitStrings {
		lengths = append(lengths, len(s))
	}
	return lengths
}

func main() {
	const (
		test1 = "oxooxoxxox"
		test2 = "xabcxdefxghi"
	)

	result1 := Pg_12_05(test1)
	result2 := Pg_12_05(test2)

	fmt.Println(result1) // Output: [1, 2, 1, 0, 1, 0]
	fmt.Println(result2) // Output: [0, 3, 3, 3]
}

// strings.Split 함수는 지정된 구분자로 문자열을 분할할 때,
// 구분자가 연속해서 나타나거나 문자열의 시작 혹은 끝에 위치하는 경우 빈 문자열("")을 결과에 포함
