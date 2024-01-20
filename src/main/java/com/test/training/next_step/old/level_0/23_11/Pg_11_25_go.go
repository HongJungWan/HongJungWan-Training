package main

import "fmt"

func main() {
	fmt.Println(Pg_11_25(12)) // Should return [2, 3]
	fmt.Println(Pg_11_25(17)) // Should return [17]
}

func Pg_11_25(n int) []int {
	var X []int
	for i := 2; i <= n; i++ {
		n, X = AppendPrimeFactor(n, i, X)
	}
	if n > 1 {
		X = append(X, n)
	}
	return X
}

func AppendPrimeFactor(n int, i int, X []int) (int, []int) {
	for n%i == 0 {
		X = AppendIfNotExists(X, i)
		n /= i
	}
	return n, X
}

func AppendIfNotExists(X []int, i int) []int {
	if len(X) == 0 || X[len(X)-1] != i {
		X = append(X, i)
	}
	return X
}
