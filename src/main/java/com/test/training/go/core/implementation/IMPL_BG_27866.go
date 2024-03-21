package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 문자열 S 입력 받기
	scanner.Scan()
	S := scanner.Text()

	// 정수 i 입력 받기
	scanner.Scan()
	iStr := scanner.Text()
	i, _ := strconv.Atoi(iStr)

	if i > 0 && i <= len(S) {
		fmt.Println(string(S[i-1]))
	}
}
