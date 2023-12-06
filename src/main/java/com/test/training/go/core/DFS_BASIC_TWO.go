package main

import (
	"fmt"
)

type VirusGraph struct {
	vertices [][]int
}

func NewVirusGraph(vertexCount int) *VirusGraph {
	return &VirusGraph{
		vertices: make([][]int, vertexCount),
	}
}

func (g *VirusGraph) AddEdge(from, to int) {
	g.vertices[from] = append(g.vertices[from], to)
	g.vertices[to] = append(g.vertices[to], from) // 양방향 연결
}

func (g *VirusGraph) CountInfectedComputers(start int) int {
	visited := make([]bool, len(g.vertices))
	count := 0
	g.dfsUtil(start, visited, &count)
	return count
}

func (g *VirusGraph) dfsUtil(vertex int, visited []bool, count *int) {
	visited[vertex] = true
	for _, v := range g.vertices[vertex] {
		if !visited[v] {
			*count++
			g.dfsUtil(v, visited, count)
		}
	}
}

func main() {
	var n, m int
	fmt.Println("Enter the number of computers and connections:")
	fmt.Scan(&n, &m)

	g := NewVirusGraph(n + 1) // 인덱스 1부터 시작

	fmt.Println("Enter the connections:")
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		g.AddEdge(a, b)
	}

	infectedCount := g.CountInfectedComputers(1) // 1번 컴퓨터에서 시작
	fmt.Printf("Number of infected computers: %d\n", infectedCount)
}

// 대표적인 깊이 우선 탐색(DFS) 문제 중 하나는 "바이러스" 문제다.
// 이 문제는 네트워크상의 컴퓨터들이 연결되어 있을 때,
// 특정 바이러스에 감염된 컴퓨터에서 시작하여 바이러스가 전파되는 컴퓨터의 수를 찾는 문제다.
//
//
// 문제 설명:
//
// 1. 컴퓨터의 수 n과 연결된 컴퓨터 쌍의 수 m이 주어진다.
// 2. 컴퓨터 쌍의 연결 정보가 주어진다.
// 3. 첫 번째 컴퓨터에서 바이러스가 시작되어 다른 컴퓨터로 전파된다.
// 4. 바이러스가 전파된 컴퓨터의 수를 찾아야 한다. (첫 번째 컴퓨터 제외)
//
//
// 5개의 컴퓨터와 4개의 연결이 있는 네트워크를 가정.
//
//
// 컴퓨터의 수 (n): 5
// 연결된 컴퓨터 쌍의 수 (m): 4
// 연결 정보:
// 1과 2가 연결
// 2와 3이 연결
// 3과 4가 연결
// 4와 5가 연결
//
//
// 입력 예시:
// 5 4
// 1 2
// 2 3
// 3 4
// 4 5
