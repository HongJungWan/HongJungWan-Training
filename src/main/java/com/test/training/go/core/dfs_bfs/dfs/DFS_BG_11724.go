package main

import (
	"fmt"
)

var (
	adjMatrix          [][]bool // 인접 행렬로 그래프 표현
	visitedNode        []bool   // 각 정점의 방문 여부
	numNodes, numEdges int      // numNodes는 정점의 개수, numEdges는 간선의 개수
)

func main() {
	fmt.Scan(&numNodes, &numEdges) // 정점과 간선의 개수를 입력

	// 인접 행렬과 방문 배열을 초기화
	adjMatrix = make([][]bool, numNodes+1)
	for i := range adjMatrix {
		adjMatrix[i] = make([]bool, numNodes+1)
	}
	visitedNode = make([]bool, numNodes+1)

	// 간선 정보를 입력받아 인접 행렬에 저장
	var node1, node2 int
	for i := 0; i < numEdges; i++ {
		fmt.Scan(&node1, &node2)
		adjMatrix[node1][node2] = true
		adjMatrix[node2][node1] = true // 무방향 그래프이므로 양방향 표시
	}

	// 연결 요소의 개수를 구함
	componentCount := 0
	for node := 1; node <= numNodes; node++ {
		if !visitedNode[node] {
			exploreNodeDFS(node)
			componentCount++ // 새 연결 요소를 찾을 때마다 카운트 증가
		}
	}

	// 연결 요소의 개수 출력
	fmt.Println(componentCount)
}

// exploreNodeDFS 함수는 주어진 노드를 시작점으로 깊이 우선 탐색을 수행
func exploreNodeDFS(node int) {
	visitedNode[node] = true // 현재 노드를 방문했음을 표시

	// 모든 노드에 대해 반복
	for adjacentNode := 1; adjacentNode <= numNodes; adjacentNode++ {
		// 현재 노드와 인접하며 아직 방문하지 않은 노드가 있다면
		if adjMatrix[node][adjacentNode] && !visitedNode[adjacentNode] {
			exploreNodeDFS(adjacentNode) // 인접 노드에서 탐색을 계속
		}
	}
}
