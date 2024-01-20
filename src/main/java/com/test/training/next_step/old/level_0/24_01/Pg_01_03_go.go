package main

import (
	"fmt"
)

func solution(myString string, m int, c int) string {
	rows := len(myString) / m
	result := ""

	for i, ch := range myString {
		row := i / m
		col := i%m + 1

		if col == c {
			result += string(ch)
			if row == rows-1 {
				break
			}
		}
	}
	return result
}

func main() {
	fmt.Println(solution("ihrhbakrfpndopljhygc", 4, 2)) // "happy"
	fmt.Println(solution("programmers", 1, 1))          // "programmers"
}
