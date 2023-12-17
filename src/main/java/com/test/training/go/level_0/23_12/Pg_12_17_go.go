package main

import (
	"fmt"
	"strconv"
)

func Pg_12_17(number string) int {
	sum := 0
	for _, digit := range number {
		d, _ := strconv.Atoi(string(digit))
		sum += d
	}
	return sum % 9
}

func main() {
	fmt.Println(Pg_12_17("123"))
	fmt.Println(Pg_12_17("78720646226947352489"))
}
