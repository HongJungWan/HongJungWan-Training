package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords) // 공백으로 구분된 단어(정수)를 읽기 위해 설정

	// N 읽기
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	// 최솟값과 최댓값 초기화
	scanner.Scan()
	firstNumber, _ := strconv.Atoi(scanner.Text())
	min, max := firstNumber, firstNumber

	for i := 1; i < N; i++ {
		scanner.Scan()
		number, _ := strconv.Atoi(scanner.Text())

		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}
	fmt.Println(min, max)
}
