package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var K int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scan(&K)

	// 숫자를 저장할 슬라이스
	stack := make([]int, 0)

	for i := 0; i < K; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())

		if num == 0 {
			// 가장 최근에 쓴 수를 지움
			stack = stack[:len(stack)-1]
		} else {
			// 수를 슬라이스에 추가
			stack = append(stack, num)
		}
	}

	sum := 0
	for _, num := range stack {
		sum += num
	}
	fmt.Println(sum)
}
