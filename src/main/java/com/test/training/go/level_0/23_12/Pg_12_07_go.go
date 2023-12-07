package main

import (
	"fmt"
	"strconv"
)

func Pg_12_07(a int, b int) int {
	concatenated, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	multiplied := 2 * a * b

	if concatenated > multiplied {
		return concatenated
	}
	return multiplied
}

func main() {
	fmt.Println(Pg_12_07(2, 91)) // 364 출력
	fmt.Println(Pg_12_07(91, 2)) // 912 출력
}

// strconv Go 언어의 문자열 변환 패키지
//
// strconv.Itoa: 이 함수는 "Integer to ASCII"의 줄임말, 문자열로 변환
// strconv.Atoi: 이 함수는 "ASCII to Integer"의 줄임말, 정수로 변환
