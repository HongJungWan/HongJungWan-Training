package main

import (
	"fmt"
	"sort"
	"strings"
)

func Pg_12_09(myString string) []string {
	splitStrings := strings.Split(myString, "x")
	var filteredStrings []string

	for _, str := range splitStrings {
		if str != "" {
			filteredStrings = append(filteredStrings, str)
		}
	}
	sort.Strings(filteredStrings)

	return filteredStrings
}

func main() {
	test1 := Pg_12_09("axbxcxdx")
	test2 := Pg_12_09("dxccxbbbxaaaa")

	fmt.Println(test1) // Output: ["a", "b", "c", "d"]
	fmt.Println(test2) // Output: ["aaaa", "bbb", "cc", "d"]
}
