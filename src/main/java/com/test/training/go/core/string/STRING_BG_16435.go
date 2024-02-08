// BG - 스네이크 버드

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, L int
	fmt.Fscanln(reader, &N, &L)

	fruits := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscanf(reader, "%d", &fruits[i])
	}
	fmt.Fscanln(reader) // 과일 높이 입력 후 버퍼 비우기

	// 과일의 높이에 따라 정렬
	sort.Ints(fruits)

	// 스네이크버드의 초기 길이를 이용하여 먹을 수 있는 과일 찾기
	for _, height := range fruits {
		if L >= height {
			L++ // 스네이크버드 길이 증가
		}
	}
	fmt.Fprintln(writer, L)
}
