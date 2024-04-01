package main

import (
	"fmt"
)

func main() {
	var n, m, person1, person2 int
	fmt.Scanf("%d\n%d %d\n%d", &n, &person1, &person2, &m)

	graph := make(map[int][]int)
	visited := make(map[int]bool)

	for i := 0; i < m; i++ {
		var parent, child int
		fmt.Scanf("%d %d", &parent, &child)
		graph[parent] = append(graph[parent], child)
		graph[child] = append(graph[child], parent) // 양방향으로 추가
	}
	result := calculateKinshipDFS(graph, visited, person1, person2, 0)
	fmt.Println(result)
}

func calculateKinshipDFS(graph map[int][]int, visited map[int]bool, start, end, depth int) int {
	if start == end {
		return depth
	}

	visited[start] = true
	for _, next := range graph[start] {
		if !visited[next] {
			distance := calculateKinshipDFS(graph, visited, next, end, depth+1)
			if distance != -1 {
				return distance
			}
		}
	}
	return -1
}
