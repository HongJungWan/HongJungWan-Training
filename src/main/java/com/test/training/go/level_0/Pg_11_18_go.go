package main

import (
	"fmt"
	"strings"
)

func Pg_11_18(my_string string, alp string) string {
	return strings.Replace(my_string, alp, strings.ToUpper(alp), -1)
}

func main() {
	fmt.Println(Pg_11_18("programmers", "p")) // "Programmers"
	fmt.Println(Pg_11_18("lowercase", "x"))   // "lowercase"
}

// func Replace(s, old, new string, n int) string
//
// s: 대상이 되는 원본 문자열.
// old: s 문자열 내에서 교체하려는 문자열.
// new: old 문자열을 대체할 새로운 문자열.

// n: 교체할 횟수를 지정.
// 	n < 0: old 문자열이 s 내에서, 발견되는 모든 곳에서 교체.
// 	n == 0: 교체가 일어나지 않는다.
// 	n > 0: old 문자열이 s 내에서, 처음 발견된 n개의 인스턴스만 교체.
