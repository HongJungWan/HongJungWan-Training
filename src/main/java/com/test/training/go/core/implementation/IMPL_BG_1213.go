package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1] // Remove newline character

	result, err := createPalindrome(name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func createPalindrome(name string) (string, error) {
	count := make([]int, 26) // 각 문자 'A'부터 'Z'까지의 빈도를 저장할 배열을 초기화

	// 입력된 문자열에 대한 각 문자의 빈도를 세어 배열에 저장
	for _, ch := range name {
		count[ch-'A']++
	}

	oddCount := 0 // 홀수 빈도의 문자 개수를 세는 변수
	oddChar := 0
	for i, c := range count {
		if c%2 == 1 {
			oddCount++
			oddChar = i // 홀수 빈도를 가진 문자의 위치(인덱스)를 저장
		}
	}

	// 조건, 홀수 빈도의 문자가 1개를 초과하면 팰린드롬을 만들 수 없다
	if oddCount > 1 {
		return "", fmt.Errorf("I'm Sorry Hansoo")
	}

	// 조건, 홀수 문자가 있을 경우 실행하는 로직
	firstHalf := make([]byte, 0, len(name)/2)
	var middleChar byte
	if oddCount == 1 {
		middleChar = byte(oddChar) + 'A' // 홀수 빈도의 문자를 중간에 위치시키기 위해 저장
		count[oddChar]--                 // 홀수 빈도 문자의 빈도를 하나 줄여 짝수로 만든다.
	}

	// 배열에 기록된 빈도수를 바탕으로 팰린드롬의 첫 부분을 생성
	for i, c := range count {
		for j := 0; j < c/2; j++ {
			firstHalf = append(firstHalf, byte(i)+'A')
		}
	}

	// 전체 팰린드롬을 구성하기 위한 배열을 생성하고, 첫 부분을 복사
	result := make([]byte, len(name))
	copy(result, firstHalf)
	if oddCount == 1 {
		result[len(firstHalf)] = middleChar // 홀수 빈도 문자를 중간에 위치
	}
	// 첫 부분의 반대 순서로 두 번째 부분을 복사하여 팰린드롬을 완성합
	for i := range firstHalf {
		result[len(name)-1-i] = firstHalf[i]
	}

	return string(result), nil
}
