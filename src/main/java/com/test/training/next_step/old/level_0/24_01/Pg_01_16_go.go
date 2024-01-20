package main

import (
	"fmt"
)

func compareDates(date1 []int, date2 []int) int {
	year1, month1, day1 := date1[0], date1[1], date1[2]
	year2, month2, day2 := date2[0], date2[1], date2[2]

	if year1 < year2 {
		return 1
	} else if year1 > year2 {
		return 0
	}

	if month1 < month2 {
		return 1
	} else if month1 > month2 {
		return 0
	}

	if day1 < day2 {
		return 1
	} else if day1 > day2 {
		return 0
	}

	return 0 // 두 날짜가 동일한 경우
}

func main() {
	date1 := []int{2021, 12, 28}
	date2 := []int{2021, 12, 29}
	result1 := compareDates(date1, date2)
	fmt.Println(result1) // 출력: 1

	date3 := []int{1024, 10, 24}
	date4 := []int{1024, 10, 24}
	result2 := compareDates(date3, date4)
	fmt.Println(result2) // 출력: 0
}
