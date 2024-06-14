// TODO: 이진 검색 트리
// TODO: 이진 트리의 최대 깊이를 구하는 문제 (DFS)
//
// TODO: 왼쪽 서브트리와 오른쪽 서브트리 중 더 큰 깊이를 선택하고, 거기에 현재 노드의 깊이(1)를 더하는 것이 이진 트리의 깊이를 구하는 기본 아이디어

package main

import "fmt"

type TreeDepth struct {
	Val   int
	Left  *TreeDepth
	Right *TreeDepth
}

func MaxDepthDFS(root *TreeDepth) int {
	if root == nil {
		return 0
	}
	leftDepth := MaxDepthDFS(root.Left)
	rightDepth := MaxDepthDFS(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func main() {
	root := &TreeDepth{Val: 3}
	root.Left = &TreeDepth{Val: 9}
	root.Right = &TreeDepth{Val: 20}
	root.Right.Left = &TreeDepth{Val: 13}
	root.Right.Right = &TreeDepth{Val: 25}

	fmt.Println("\nMax Depth:", MaxDepthDFS(root)) // Output: Max Depth: 3
}

/*
     3
    / \
   9  20
     /  \
    13   25
*/

/*
MaxDepthDFS(root *TreeDepth) int 호출

1. [root] MaxDepthDFS(3)
   - root.Left와 root.Right 재귀 호출

2. [root.Left] MaxDepthDFS(9)
   - leftDepth = 1 (노드 9)

3. [root.Right] MaxDepthDFS(20)
   - root.Right.Left와 root.Right.Right 재귀 호출

4. [root.Right.Left] MaxDepthDFS(13)
   - rightLeftDepth = 1 (노드 13)

5. [root.Right.Right] MaxDepthDFS(25)
   - rightRightDepth = 1 (노드 25)

6. rightDepth = 1 + max(1, 1) = 2 (노드 20)

7. root의 최대 깊이 = 1 + max(leftDepth, rightDepth)
   - leftDepth = 1
   - rightDepth = 2
   - 결과: 3
*/
