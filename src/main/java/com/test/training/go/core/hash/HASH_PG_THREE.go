// PG - 전화번호 목록
package main

import (
	"sort"
	"strings"
)

func isUniquePrefix(phoneBook []string) bool {
	sort.Strings(phoneBook)

	for i := 0; i < len(phoneBook)-1; i++ {
		if strings.HasPrefix(phoneBook[i+1], phoneBook[i]) {
			return false // 접두어인 경우, false 반환.
		}
	}
	return true
}

func main() {
	phoneBook1 := []string{"119", "97674223", "1195524421"}
	phoneBook2 := []string{"123", "456", "789"}
	phoneBook3 := []string{"12", "123", "1235", "567", "88"}

	println(isUniquePrefix(phoneBook1)) // false
	println(isUniquePrefix(phoneBook2)) // true
	println(isUniquePrefix(phoneBook3)) // false
}
