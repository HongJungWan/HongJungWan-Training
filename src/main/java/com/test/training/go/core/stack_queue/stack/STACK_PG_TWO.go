// PG - 올바른 괄호

package main

import (
	"fmt"
)

func isValidParentheses(s string) bool {
	stack := []rune{}

	for _, char := range s {
		switch char {
		case '(':
			stack = append(stack, char)
		case ')':
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValidParentheses("()()"))   // true
	fmt.Println(isValidParentheses("(())()")) // true
	fmt.Println(isValidParentheses(")()("))   // false
	fmt.Println(isValidParentheses("(()("))   // false
}
