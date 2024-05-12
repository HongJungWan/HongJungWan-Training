package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var N int
	fmt.Sscan(scanner.Text(), &N)

	weight := make(map[rune]int)

	// 각 단어를 스캔하면서 각 글자의 자릿수 가치를 계산
	for i := 0; i < N; i++ {
		scanner.Scan()
		word := scanner.Text()
		l := len(word)
		for j, ch := range word {
			// 10의 거듭제곱 자릿수 가치 계산
			weight[ch] += pow(10, l-j-1)
		}
	}

	// 자릿수 가치를 기준으로 정렬하기 위해 슬라이스로 변환
	weights := make([]int, 0, len(weight))
	for _, w := range weight {
		weights = append(weights, w)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(weights)))

	// 가장 큰 자릿수 가치부터 9, 8, ..., 1, 0을 할당
	result := 0
	num := 9
	for _, w := range weights {
		result += w * num
		num--
	}

	fmt.Println(result)
}

// 거듭제곱 계산 함수
func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

/*
1. 자릿수 가치 계산

1-1. 각 단어에 대해 반복하면서, 각 문자가 위치한 자릿수에 따라 그 가치를 계산.
예를 들어, 단어 "ABC"에서 'A', 'B', 'C'는 각각 백의 자리, 십의 자리, 일의 자리에 있으므로 가치는 100, 10, 1

1-2. 각 문자의 자릿수 가치를 누적하여 저장.
예를 들어, 'A'가 여러 단어에서 백의 자리에 있다면 그 가치는 여러 번 더해집니다.


2. 가치에 따라 숫자 할당

계산된 자릿수 가치를 기준으로 내림차순 정렬.
가장 높은 가치를 가진 문자부터 숫자 9, 8, ..., 0을 할당.


3. 할당된 숫자를 각 문자의 자릿수 가치에 곱한 후, 모두 더한다.
*/
