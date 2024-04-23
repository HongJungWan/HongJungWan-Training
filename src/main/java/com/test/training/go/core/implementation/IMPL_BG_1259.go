package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isPalindromes(num int) bool {
	s := strconv.Itoa(num) // 정수를 문자열로 변환
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		number, _ := strconv.Atoi(input)
		if number == 0 {
			break
		}

		if isPalindromes(number) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
