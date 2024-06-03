// TODO: 문자열 압축 문제
// TODO: 연속된 동일 문자를, 문자와 그 개수로 대체 -> 압축

package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) {
	n := len(chars)

	index := 0 // 압축된 문자열을 쓰기 위한 위치
	count := 1

	for i := 1; i <= n; i++ {
		// 현재 문자와 이전 문자가 동일하면 count 증가
		if i < n && chars[i] == chars[i-1] {
			count++
		} else {
			// 현재 문자를 배열에 저장한 후 인덱스 증가
			chars[index] = chars[i-1]
			index++

			// count가 1보다 크면 count를 문자열로 변환하여 배열에 작성
			if count > 1 {
				countStr := strconv.Itoa(count)
				for _, c := range countStr {
					chars[index] = byte(c)
					index++ // 문자의 개수를 배열에 작성한 후 인덱스를 증가시키기 위해 사용
				}
			}

			// 새로운 문자에 대해 count를 1로 초기화
			count = 1
		}
	}
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	compress(chars)
	fmt.Printf("\nCompressed array: %s\n", chars)
}
