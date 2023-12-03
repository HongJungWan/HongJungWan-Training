package main

import (
	"fmt"
)

func Pg_12_02(myStrings []string, parts [][]int) string {
	result := ""
	for i, part := range parts {
		s, e := part[0], part[1]
		result += myStrings[i][s : e+1]
	}
	return result
}

func main() {
	myStrings := []string{"progressive", "hamburger", "hammer", "ahocorasick"}
	parts := [][]int{{0, 4}, {1, 2}, {3, 5}, {7, 7}}
	fmt.Println(Pg_12_02(myStrings, parts))
}
