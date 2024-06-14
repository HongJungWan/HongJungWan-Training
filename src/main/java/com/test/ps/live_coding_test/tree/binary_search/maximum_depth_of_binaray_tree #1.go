// TODO: 이진 검색 트리
// TODO: 이진 트리의 최대 깊이를 구하는 문제 (DFS)

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
    15   7
*/

/*

 */
