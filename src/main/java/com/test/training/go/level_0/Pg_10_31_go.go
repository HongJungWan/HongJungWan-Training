package main

import "fmt"

func solution(n int) int {
	result := 0

	if n%2 == 1 {
		for i := 1; i <= n; i += 2 {
			result += i
		}
	} else {
		for i := 2; i <= n; i += 2 {
			result += i * i
		}
	}

	return result
}

func main() {
	fmt.Println(solution(7))  // 16
	fmt.Println(solution(10)) // 220
}