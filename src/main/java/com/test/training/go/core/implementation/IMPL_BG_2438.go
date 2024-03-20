package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N) // 사용자로부터 N의 값을 입력받음

	for i := 1; i <= N; i++ { // 1부터 N까지 반복
		for j := 0; j < i; j++ { // 각 줄에 i개의 별을 출력
			fmt.Print("*")
		}
		fmt.Println()
	}
}
