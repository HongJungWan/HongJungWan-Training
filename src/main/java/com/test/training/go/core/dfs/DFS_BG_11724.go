package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	adjList            [][]int // 인접 리스트로 그래프 표현
	visitedNode        []bool  // 각 정점의 방문 여부
	numNodes, numEdges int     // 정점의 개수와 간선의 개수
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &numNodes, &numEdges)

	// 인접 리스트와 방문 배열을 초기화
	adjList = make([][]int, numNodes+1)
	visitedNode = make([]bool, numNodes+1)

	// 간선 정보를 입력받아 인접 리스트에 저장
	var node1, node2 int
	for i := 0; i < numEdges; i++ {
		fmt.Fscan(reader, &node1, &node2)
		adjList[node1] = append(adjList[node1], node2)
		adjList[node2] = append(adjList[node2], node1) // 무방향 그래프이므로 양방향 추가
	}

	// 연결 요소의 개수를 구함
	componentCount := 0
	for node := 1; node <= numNodes; node++ {
		if !visitedNode[node] {
			exploreNodeDFS(node)
			componentCount++
		}
	}

	// 연결 요소의 개수 출력
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(writer, componentCount)
	writer.Flush()
}

// exploreNodeDFS 함수는 주어진 노드를 시작점으로 깊이 우선 탐색을 수행
func exploreNodeDFS(node int) {
	visitedNode[node] = true

	// 인접 리스트를 사용하여 인접 노드 반복
	for _, adjacentNode := range adjList[node] {
		if !visitedNode[adjacentNode] {
			exploreNodeDFS(adjacentNode)
		}
	}
}

// 정점의 수가 적을 때 -> 인접 행렬
// 인접 행렬을 사용할 경우, 간선의 조회에 O(n^2)의 시간이 소요된다. 이는 각 정점 쌍에 대해 연결 여부를 확인해야 하기 때문이다.
// 정점의 수(n)가 많을 때, 이 방법은 매우 비효율적이 될 수 있다.
// 특히, 간선의 수(e)가 정점의 수에 비해 상대적으로 적은 "희소 그래프(sparse graph)"의 경우 이러한 비효율성은 더욱 두드러진다.

// 정점의 수가 많을 때 -> 인접 리스트
// 인접 리스트를 사용하는 방법은 이러한 문제를 해결할 수 있는 방법이다.
// 인접 리스트는 각 정점에서 직접 연결된 간선들만을 리스트로 관리하여, 간선의 조회 시간을 O(V+E)로 줄일 수 있다.
// 여기서 V는 정점의 수, E는 간선의 수를 의미한다.
// 이 방식은 각 정점에서 출발하는 간선들을 바로 접근할 수 있기 때문에, 그래프 탐색과 같은 작업을 보다 효율적으로 수행할 수 있다.
// 특히, 희소 그래프에서 각 정점에 연결된 간선의 수가 적기 때문에, 전체 작업의 시간 복잡도가 상당히 감소한다.
