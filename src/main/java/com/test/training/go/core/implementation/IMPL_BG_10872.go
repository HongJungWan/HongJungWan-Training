package main

import (
	"fmt"
)

// factorial 함수는 재귀적으로 팩토리얼을 계산합니다.
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var n int
	fmt.Scan(&n)              // 사용자로부터 정수 N을 입력받습니다.
	fmt.Println(factorial(n)) // 팩토리얼을 계산하여 출력합니다.
}
