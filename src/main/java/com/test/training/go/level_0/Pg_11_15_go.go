package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Pg_11_15(myString string) string {
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
	fmt.Println(Pg_11_15("abstract algebra")) // "AbstrAct AlgebrA"
	fmt.Println(Pg_11_15("PrOgRaMmErS"))      // "progrAmmers"
}
