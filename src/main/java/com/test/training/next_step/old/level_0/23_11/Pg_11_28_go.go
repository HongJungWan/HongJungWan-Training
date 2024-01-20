package main

import (
	"fmt"
)

func Pg_11_28(names []string) []string {
	var result []string

	for i := 0; i < len(names); i += 5 {
		result = append(result, names[i])
	}
	return result
}

func main() {
	names := []string{"nami", "ahri", "jayce", "garen", "ivern", "vex", "jinx"}
	fmt.Println(Pg_11_28(names))
}
