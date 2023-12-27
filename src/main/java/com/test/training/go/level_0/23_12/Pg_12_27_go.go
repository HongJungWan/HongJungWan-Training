package main

import (
	"fmt"
	"sort"
)

func Pg_12_27(my_string string) []string {
	length := len(my_string)
	suffixes := make([]string, length)

	for i := 0; i < length; i++ {
		suffixes[i] = my_string[i:]
	}
	sort.Strings(suffixes)
	return suffixes
}

func main() {
	fmt.Println(Pg_12_27("banana"))      // ["a", "ana", "anana", "banana", "na", "nana"]
	fmt.Println(Pg_12_27("programmers")) // ["ammers", "ers", "grammers", "mers", "mmers", "ogrammers", "programmers", "rammers", "rogrammers", "rs", "s"]
}
