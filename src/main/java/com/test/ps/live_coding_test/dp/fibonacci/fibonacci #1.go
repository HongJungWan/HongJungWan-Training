// TODO: 피보나치
// TODO: 피보나치 수열에서 다음 수는 앞의 두 수를 더한 값으로 이루어지는 수열

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
	fibonacciNumbers := fibonacci(n)

	fmt.Println("Fibonacci sequence:", fibonacciNumbers)
}

func fibonacci(n int) []int {
	if n < 2 {
		return []int{1}
	}

	fibNums := make([]int, n)
	fibNums[0], fibNums[1] = 1, 1

	for i := 2; i < n; i++ {
		fibNums[i] = fibNums[i-1] + fibNums[i-2]
	}
	return fibNums
}
