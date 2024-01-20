package main

import (
	"fmt"
	"strings"
)

func AppendToSequence(numLog []int) string {
	var sb strings.Builder

	for i := 0; i < len(numLog)-1; i++ {
		diff := numLog[i+1] - numLog[i]

		if diff == 1 {
			sb.WriteString("w")
		} else if diff == -1 {
			sb.WriteString("s")
		} else if diff == 10 {
			sb.WriteString("d")
		} else if diff == -10 {
			sb.WriteString("a")
		}
	}
	return sb.String()
}

func main() {
	numLog := []int{0, 1, 0, 10, 0, 1, 0, 10, 0, -1, -2, -1}
	fmt.Print(AppendToSequence(numLog))
}
