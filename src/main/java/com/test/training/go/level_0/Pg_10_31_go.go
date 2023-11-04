package main

import "fmt"

func Pg_10_31(n int) int {
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
	fmt.Println(Pg_10_31(7))  // 16
	fmt.Println(Pg_10_31(10)) // 220
}
