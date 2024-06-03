// TODO: DFS (기초)
// TODO: 인접 리스트 방식으로 구현, 재귀를 사용한 방식

package main

import "fmt"

func DFS(graph map[int][]int, start int, visited map[int]bool) {
	// 현재 노드를 방문 처리
	visited[start] = true
	fmt.Println(start)

	// 인접 노드들을 방문
	for _, node := range graph[start] {
		if !visited[node] {
			DFS(graph, node, visited)
		}
	}
}

func main() {
	// 그래프를 인접 리스트로 표현
	graph := map[int][]int{
		0: {1, 2},
		1: {2},
		2: {0, 3},
		3: {3},
	}

	// 방문 여부를 기록하는 맵
	visited := make(map[int]bool)

	// DFS 탐색 시작
	DFS(graph, 2, visited)
}

/*
인접 리스트: 간선이 노드 수에 비해 적은 경우 사용 (희소 그래프)
인접 행렬: 간선이 노드 수에 비해 많은 경우 사용 (밀집 그래프)
*/
