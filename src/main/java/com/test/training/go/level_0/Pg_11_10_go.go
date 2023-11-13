package main

import (
	"fmt"
	"strings"
)

func Pg_11_10(n_str string) string {
	// strings.TrimLeft 함수를 사용하여 왼쪽의 "0"들을 제거한다.
	return strings.TrimLeft(n_str, "0")
}

func main() {
	fmt.Println(Pg_11_10("0010"))   // "10"
	fmt.Println(Pg_11_10("854020")) // "854020"
}
