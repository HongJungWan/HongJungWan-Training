package main

import (
	"fmt"
	"strings"
	"unicode"
)

func solution(myString string) string {
	var result strings.Builder

	for _, ch := range myString {
		if ch == 'a' {
			result.WriteRune('A')
		} else if unicode.IsUpper(ch) && ch != 'A' {
			result.WriteRune(unicode.ToLower(ch))
		} else {
			result.WriteRune(ch)
		}
	}

	return result.String()
}

func main() {
	fmt.Println(solution("abstract algebra")) // "AbstrAct AlgebrA"
	fmt.Println(solution("PrOgRaMmErS"))      // "progrAmmers"
}
