// TODO: BFS (기초)
// TODO: 인접 리스트 방식으로 구현, 큐를 사용한 방식

package main

import (
	"container/list"
	"fmt"
)

func BFS(graph map[int][]int, start int, visited map[int]bool) []int {
	// 큐 생성
	queue := list.New()
	result := []int{}

	// 시작 노드를 큐에 추가하고 방문 처리
	queue.PushBack(start)
	visited[start] = true

	// 큐가 빌 때까지 반복
	for queue.Len() > 0 {
		// 큐의 앞에서 노드를 꺼냅니다.
		element := queue.Front()
		node := element.Value.(int)
		queue.Remove(element)
		result = append(result, node)

		// 현재 노드의 이웃 노드들을 순회
		for _, neighbor := range graph[node] {

			// 방문하지 않은 이웃 노드를 큐에 추가하고 방문 처리
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
			}
		}
	}
	return result
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

	// BFS 탐색 시작
	result := BFS(graph, 2, visited)

	// 방문한 노드들을 출력
	fmt.Println("BFS 방문 순서:", result)
}

/*
인접 리스트: 간선이 노드 수에 비해 적은 경우 사용 (희소 그래프)
인접 행렬: 간선이 노드 수에 비해 많은 경우 사용 (밀집 그래프)


container/list 패키지는 Go 표준 라이브러리 중 하나로, 양방향 연결 리스트(Doubly Linked List)를 제공하는 패키지

queue := list.New(), list.New()를 호출하여 새로운 양방향 연결 리스트를 생성
queue.PushBack(start), PushBack 메서드를 사용하여 큐의 끝부분에 새로운 요소를 추가

element := queue.Front(), Front 메서드를 사용하여 큐의 앞에 있는 요소를 가져오고
node := element.Value.(int), 그 값을 Value 필드를 통해 가져와 타입 변환
queue.Remove(element), 그 후 Remove 메서드를 사용하여 큐에서 해당 요소를 제거
*/
