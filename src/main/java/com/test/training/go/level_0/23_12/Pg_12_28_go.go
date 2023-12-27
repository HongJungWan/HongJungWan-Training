package main

import "fmt"

func Pg_12_28(n int) []int {
	number := []int{n}

	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		number = append(number, n)
	}
	return number
}

func main() {
	n := 10
	fmt.Println(Pg_12_28(n))
}
