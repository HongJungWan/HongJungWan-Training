package main

import (
	"fmt"
)

func isLeapYear(year int) int {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return 1 // 윤년일 경우
	}
	return 0 // 윤년이 아닐 경우
}

func main() {
	var year int
	fmt.Scan(&year)

	fmt.Println(isLeapYear(year))
}
