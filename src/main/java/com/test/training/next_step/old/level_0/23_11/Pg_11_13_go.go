package main

import (
	"fmt"
)

func Pg_11_13(my_string string, is_suffix string) int {
	if len(is_suffix) > len(my_string) {
		return 0
	}
	// my_string의 끝에서부터 is_suffix의 길이만큼의 부분을 가져온다.
	endPart := my_string[len(my_string)-len(is_suffix):]

	if endPart == is_suffix {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(Pg_11_13("banana", "ana"))     // 1
	fmt.Println(Pg_11_13("banana", "nan"))     // 0
	fmt.Println(Pg_11_13("banana", "wxyz"))    // 0
	fmt.Println(Pg_11_13("banana", "abanana")) // 0
}
