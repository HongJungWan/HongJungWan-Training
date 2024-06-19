// TODO: 다익스트라 알고리즘
// TODO: 주요 개념 (bFS, Graph, Heap[Priority Queue], Shortest Path)

package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to, time int
}

type MinHeap []Node

type Node struct {
	id, dist int
}

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	times := [][]int{
		{2, 1, 1},
		{2, 3, 1},
		{3, 4, 1},
	}
	n := 4
	k := 2
	result := networkDelayTime(times, n, k)
	fmt.Println(result) // Output: 2
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph, dist, minHeap := setupGraphAndDistances(times, n, k)

	dijkstraAlgorithm(minHeap, dist, graph)

	maxDist, i, done := calculateNetworkDelay(n, dist)
	if done {
		return i
	}

	return maxDist
}

func setupGraphAndDistances(times [][]int, n int, k int) ([][]Edge, []int, *MinHeap) {
	// 그래프 초기화
	graph := make([][]Edge, n+1)
	for _, time := range times {
		u, v, w := time[0], time[1], time[2]
		graph[u] = append(graph[u], Edge{to: v, time: w})
	}

	// 거리 배열 초기화
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0

	// 최소 힙 초기화
	minHeap := &MinHeap{}
	heap.Push(minHeap, Node{id: k, dist: 0})
	return graph, dist, minHeap
}

func dijkstraAlgorithm(minHeap *MinHeap, dist []int, graph [][]Edge) {
	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(Node)
		u := node.id
		d := node.dist

		if d > dist[u] {
			continue
		}

		for _, edge := range graph[u] {
			v, w := edge.to, edge.time
			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				heap.Push(minHeap, Node{id: v, dist: dist[v]})
			}
		}
	}
}

// 다익스트라 알고리즘이 종료된 후, 모든 노드에 도달 가능한지 확인하고, 도달 가능하다면 그 중 가장 긴 시간이 네트워크 지연 시간으로 계산되는 과정을 수행
func calculateNetworkDelay(n int, dist []int) (int, int, bool) {
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return 0, -1, true
		}
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}
	return maxDist, 0, false
}
