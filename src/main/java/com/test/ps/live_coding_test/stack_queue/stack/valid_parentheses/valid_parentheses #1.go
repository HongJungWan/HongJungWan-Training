// TODO: 괄호 문제
// TODO:

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	for i := 0; i < N; i++ {
		var inputParentheses string
		fmt.Fscanf(reader, "%s\n", &inputParentheses)

		if isValidParentheses(inputParentheses) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func isValidParentheses(s string) bool {
	stack := []rune{}
	for _, char := range s {
		if char == '(' {
			stack = append(stack, char)

		} else if char == ')' {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
