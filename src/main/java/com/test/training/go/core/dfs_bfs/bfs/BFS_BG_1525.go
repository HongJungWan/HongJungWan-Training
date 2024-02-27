// BG - 퍼즐

package main

import (
	"fmt"
	"strconv"
)

// 퍼즐의 상태를 표현하는 구조체
type State struct {
	board     string // 현재 보드 상태를 문자열로 표현 (0은 빈 칸)
	zeroPos   int    // 0(빈 칸)의 위치
	moveCount int    // 이동 횟수
}

// 상, 하, 좌, 우 이동에 따른 인덱스 변화
var dirs = []int{-3, 3, -1, 1}

func bfs(initialState State) int {
	visited := make(map[string]bool) // 방문한 상태를 기록
	queue := []State{initialState}   // BFS를 위한 큐

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// 목표 상태에 도달한 경우 이동 횟수 반환
		if current.board == "123456780" {
			return current.moveCount
		}

		// 이미 방문한 상태면 스킵
		if visited[current.board] {
			continue
		}
		visited[current.board] = true

		// 가능한 모든 이동을 탐색
		for _, dir := range dirs {
			nextZeroPos := current.zeroPos + dir
			if nextZeroPos >= 0 && nextZeroPos < 9 && (dir != 1 || current.zeroPos%3 != 2) && (dir != -1 || current.zeroPos%3 != 0) {
				// 이동 가능한 경우 새로운 상태 생성
				newBoard := []rune(current.board)
				newBoard[current.zeroPos], newBoard[nextZeroPos] = newBoard[nextZeroPos], newBoard[current.zeroPos]
				newState := State{string(newBoard), nextZeroPos, current.moveCount + 1}
				queue = append(queue, newState)
			}
		}
	}
	// 이동이 불가능한 경우
	return -1
}

func main() {
	var initialBoard [9]int
	var zeroPos int
	for i := 0; i < 9; i++ {
		fmt.Scan(&initialBoard[i])
		if initialBoard[i] == 0 {
			zeroPos = i
		}
	}

	// 초기 상태를 문자열로 변환
	boardStr := ""
	for _, num := range initialBoard {
		boardStr += strconv.Itoa(num)
	}

	initialState := State{boardStr, zeroPos, 0}
	fmt.Println(bfs(initialState))
}
