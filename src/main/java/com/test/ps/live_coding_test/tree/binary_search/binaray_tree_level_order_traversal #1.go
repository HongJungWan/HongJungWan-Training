// TODO: 이진 탐색 트리
// TODO: 이진 탐색 트리 레벨 순회 (BFS)

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBFS(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	var queue []*TreeNode

	// BFS 시작: 루트를 큐에 추가
	queue = append(queue, root)

	for len(queue) > 0 {
		levelSize := len(queue)
		var level []int

		// 현재 레벨의 모든 노드를 처리
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:] // 큐에서 첫 번째 노드 제거
			level = append(level, node.Val)

			// 왼쪽 자식이 있으면 큐에 추가
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			// 오른쪽 자식이 있으면 큐에 추가
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 현재 레벨의 값을 결과에 추가
		result = append(result, level)
	}

	return result
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	// BFS를 이용한 레벨 순서 탐색
	result := levelOrderBFS(root)
	fmt.Println(result)
}

/*
     3
    / \
   9  20
     /  \
    15   7
*/

/*
이진 트리의 경우와 일반적인 그래프의 BFS 구현에서 방문 처리가 다른 이유는 데이터 구조와 탐색 방식의 차이

이진 트리:
이진 트리의 특성상 노드가 중복되지 않고, 모든 노드는 최대 하나의 부모 노드를 가진다.
트리에서 BFS를 수행할 때, 각 노드를 큐에 넣고 한 번만 방문하게 된다. 즉, 이미 큐에 넣은 노드는 다시 큐에 넣지 않으므로 중복 방문이 발생하지 않는다.

[일반 그래프]
일반 그래프는 사이클(순환)을 가질 수 있으며, 노드 간의 연결 관계가 다양.
같은 노드가 여러 경로를 통해 다시 큐에 추가될 수 있기 때문에, 노드 방문 여부를 확인하고 중복 방문을 방지하기 위해 visited 맵을 사용.
*/

/*
이진 트리: 노드 중복이 없고 사이클이 없기 때문에 방문 처리를 생략해도 된다.
일반 그래프: 사이클과 다중 경로로 인해 동일 노드가 여러 번 큐에 추가될 수 있으므로 방문 처리가 필요.

즉, 이진 트리에서는 방문 처리를 생략할 수 있지만, 일반 그래프에서는 방문 처리를 반드시 해야 한다.
*/

/*
[첫 번째 while 루프 (첫 번째 레벨)]

1. queue의 길이를 levelSize에 저장 (levelSize = 1)

2. for 루프 시작 (levelSize 만큼 반복)

2-1. queue의 첫 번째 노드를 node에 저장 (node = 3)
2-1. node의 값을 level에 추가 (level = [3])
2-1. node의 왼쪽 자식 9를 queue에 추가
2-1. node의 오른쪽 자식 20을 queue에 추가

3. level 슬라이스를 result에 추가 (result = [[3]])

---

[두 번째 while 루프 (두 번째 레벨)]

1. queue의 길이를 levelSize에 저장 (levelSize = 2)

2. for 루프 시작 (levelSize 만큼 반복)

2-1-1. 첫 번째 반복 (i = 0)
2-1-2. queue의 첫 번째 노드를 node에 저장 (node = 9)
2-1-3. node의 값을 level에 추가 (level = [9])
2-1-4. node는 자식이 없으므로 queue에 추가 없음

2-2-1. 두 번째 반복 (i = 1)
2-2-2. queue의 첫 번째 노드를 node에 저장 (node = 20)
2-2-3.node의 값을 level에 추가 (level = [9, 20])
2-2-4.node의 왼쪽 자식 15를 queue에 추가
2-2-5.node의 오른쪽 자식 7을 queue에 추가

3. level 슬라이스를 result에 추가 (result = [[3], [9, 20]])

---

...
*/
