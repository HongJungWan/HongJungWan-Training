// TODO: 팰린드롬 체크 문제
// TODO: 어떤 단어를 뒤에서부터 읽어도 똑같다면 그 단어를 팰린드롬

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		if input == "0" {
			break
		}
		if isPalindrome(input) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func isPalindrome(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}
