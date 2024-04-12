package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]

	if isPalindrome(input) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
