package main

import (
	"fmt"
	"strconv"
)

func Pg_11_19(a int, b int) int {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)

	ab := aStr + bStr
	ba := bStr + aStr

	abInt, _ := strconv.Atoi(ab)
	baInt, _ := strconv.Atoi(ba)

	if abInt > baInt {
		return abInt
	}
	return baInt
}

func main() {
	fmt.Println(Pg_11_19(9, 91)) // return 991
	fmt.Println(Pg_11_19(89, 8)) // return 898
}

// Itoa는 "Integer to ASCII"의 줄임말
// 정수(int)를 문자열로 변환

// Atoi는 "ASCII to Integer"의 줄임말
// 문자열을 받아 정수(int)로 변환
