package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(calculateASONSize(input))
}

// calculateASONSize 함수는 전체 ASON 객체의 크기를 계산합
func calculateASONSize(ason string) int {
	size := 8 // 시작 ASON 객체의 기본 크기
	i := 0    // 인덱스 초기화

	for i < len(ason) {
		if ason[i] == '[' {
			// 새 ASON 객체가 시작되었을 때
			size += 8
			i++
		} else if ason[i] == ']' {
			// ASON 객체가 끝났을 때
			i++
		} else {
			// 다음 원소 찾기
			element, nextIndex := getNextElement(ason, i)
			size += calculateElementSize(element)
			i = nextIndex
		}
	}
	return size
}

// getNextElement 함수는 주어진 문자열에서 다음 원소와 그 이후 인덱스를 반환
func getNextElement(ason string, start int) (element string, nextIndex int) {
	for nextIndex = start; nextIndex < len(ason); nextIndex++ {
		if unicode.IsSpace(rune(ason[nextIndex])) {
			if nextIndex > start {
				break
			}
			continue
		}
	}
	element = ason[start:nextIndex]
	return element, nextIndex
}

// calculateElementSize 함수는 원소의 유형에 따라 크기를 계산
func calculateElementSize(element string) int {
	if len(element) == 0 { // 빈 원소
		return 0
	}
	if unicode.IsDigit(rune(element[0])) { // 양의 정수 자료형
		return 8
	}
	if unicode.IsLetter(rune(element[0])) { // 문자열 자료형
		return len(element) + 12
	}
	return 0 // 다른 유형의 원소는 0으로 가정
}
