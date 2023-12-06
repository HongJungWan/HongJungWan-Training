package main

import "fmt"

// Graph는 그래프를 나타내며, 각 노드의 연결 리스트를 포함.
type Graph struct {
	vertices [][]int
}

// NewGraph는 새로운 그래프를 생성.
func NewGraph(vertexCount int) *Graph {
	return &Graph{
		vertices: make([][]int, vertexCount),
	}
}

// AddEdge는 그래프에 새로운 간선을 추가.
func (g *Graph) AddEdge(from, to int) {
	g.vertices[from] = append(g.vertices[from], to)
}

// DFS는 주어진 노드에서 시작하여 깊이 우선 탐색을 수행.
func (g *Graph) DFS(start int) {
	visited := make([]bool, len(g.vertices))
	g.dfsUtil(start, visited)
}

// dfsUtil는 실제 깊이 우선 탐색을 수행하는 재귀 함수.
func (g *Graph) dfsUtil(vertex int, visited []bool) {
	visited[vertex] = true
	fmt.Printf("%d ", vertex)

	for _, v := range g.vertices[vertex] {
		if !visited[v] {
			g.dfsUtil(v, visited)
		}
	}
}

func main() {
	g := NewGraph(4)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	fmt.Println("DFS starting from vertex 2:")
	g.DFS(2)
}
