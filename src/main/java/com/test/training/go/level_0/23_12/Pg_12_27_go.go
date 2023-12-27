package main

import (
	"fmt"
	"sort"
)

func solution(my_string string) []string {
	length := len(my_string)
	suffixes := make([]string, length)

	for i := 0; i < length; i++ {
		suffixes[i] = my_string[i:]
	}
	sort.Strings(suffixes)
	return suffixes
}

func main() {
	fmt.Println(solution("banana"))      // ["a", "ana", "anana", "banana", "na", "nana"]
	fmt.Println(solution("programmers")) // ["ammers", "ers", "grammers", "mers", "mmers", "ogrammers", "programmers", "rammers", "rogrammers", "rs", "s"]
}
