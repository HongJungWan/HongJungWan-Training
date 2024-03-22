package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	original := n
	count := 0

	for {
		n = (n%10)*10 + ((n/10)+(n%10))%10
		count++

		if n == original {
			break
		}
	}
	fmt.Println(count)
}
