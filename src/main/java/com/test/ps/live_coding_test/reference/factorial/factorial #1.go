// TODO: 팩토리얼
// TODO: 팩토리얼은 주어진 숫자부터 1까지의 모든 정수의 곱

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	n, _ := strconv.Atoi(input)

	result := factorial(n)
	fmt.Printf("Factorial of %d is %d\n", n, result)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
