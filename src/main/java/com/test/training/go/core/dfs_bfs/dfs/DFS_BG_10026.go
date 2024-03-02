package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	input   int
	grid    []string
	visited [][]bool
	dx      = []int{0, 0, 1, -1}
	dy      = []int{1, -1, 0, 0}
)

func colorBlindDFS(x, y int, c byte, colorBlind bool) {
	visited[x][y] = true
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx >= 0 && nx < input && ny >= 0 && ny < input && !visited[nx][ny] {
			if colorBlind {
				if (c == 'R' || c == 'G') && (grid[nx][ny] == 'R' || grid[nx][ny] == 'G') {
					colorBlindDFS(nx, ny, c, colorBlind)
				} else if c == grid[nx][ny] {
					colorBlindDFS(nx, ny, c, colorBlind)
				}
			} else if grid[nx][ny] == c {
				colorBlindDFS(nx, ny, c, colorBlind)
			}
		}
	}
}

func countRegions(colorBlind bool) int {
	count := 0
	visited = make([][]bool, input)
	for i := range visited {
		visited[i] = make([]bool, input)
	}
	for i := 0; i < input; i++ {
		for j := 0; j < input; j++ {
			if !visited[i][j] {
				colorBlindDFS(i, j, grid[i][j], colorBlind)
				count++
			}
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &input)
	grid = make([]string, input)
	for i := 0; i < input; i++ {
		fmt.Fscan(reader, &grid[i])
	}

	fmt.Println(countRegions(false), countRegions(true))
}
