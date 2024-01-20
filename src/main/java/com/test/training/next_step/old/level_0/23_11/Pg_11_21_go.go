package main

import (
	"fmt"
	"math"
)

func Pg_11_21(a int, b int) int {
	if a%2 != 0 && b%2 != 0 {
		return a*a + b*b
	} else if (a%2 != 0 && b%2 == 0) || (a%2 == 0 && b%2 != 0) {
		return 2 * (a + b)
	} else {
		return int(math.Abs(float64(a - b)))
	}
}

func main() {
	fmt.Println(Pg_11_21(3, 5)) // 34
	fmt.Println(Pg_11_21(6, 1)) // 14
	fmt.Println(Pg_11_21(2, 4)) // 2
}
