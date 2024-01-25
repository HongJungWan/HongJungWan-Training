// BG - 소트인사이드

package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortDescending(N string) {
	// 문자열을 문자 슬라이스로 변환
	digits := strings.Split(N, "")

	// 내림차순으로 정렬
	sort.Sort(sort.Reverse(sort.StringSlice(digits)))

	// 정렬된 문자를 결합하여 출력
	fmt.Println(strings.Join(digits, ""))
}

func main() {
	var N string
	fmt.Scan(&N)

	SortDescending(N)
}

// [Go 언어 표준 라이브러리]

// strings.Split() : 문자열을 지정된 구분자(delimiter)에 따라 분리하여 슬라이스(slice)로 반환
//
// s := "a,b,c"
// parts := strings.Split(s, ",")
// fmt.Println(parts) // ["a", "b", "c"]

// sort.Sort() : sort.Interface를 구현하는 어떤 타입의 슬라이스도 정렬
//
// numbers := []int{3, 1, 4, 1}
// sort.Sort(sort.IntSlice(numbers))
// fmt.Println(numbers) // [1, 1, 3, 4], 오름차순 정렬

// sort.Reverse() : 주어진 sort.Interface의 순서를 반전시키는 sort.Interface를 반환
//
// numbers := []int{3, 1, 4, 1}
// sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
// fmt.Println(numbers) // [4, 3, 1, 1], 내림차순 정렬

// sort.StringSlice() : 문자열 슬라이스를 정렬하는 데 사용
//
// words := []string{"peach", "banana", "apple"}
// sort.Sort(sort.StringSlice(words))
// fmt.Println(words) // ["apple", "banana", "peach"]

// strings.Join() : 스트링 슬라이스의 요소들을 지정된 구분자로 결합
//
// parts := []string{"a", "b", "c"}
// s := strings.Join(parts, ",")
// fmt.Println(s) // "a,b,c"
