// TODO: 이진 탐색 트리
// TODO: 유효한 이진 탐색 트리 확인 (DFS)

package main

import (
	"fmt"
	"math"
)

type isTreeValidNode struct {
	Val   int
	Left  *isTreeValidNode
	Right *isTreeValidNode
}

// 유효한 이진 탐색 트리인지 확인하는 함수
func isValidBST(root *isTreeValidNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

// 유효성 검사를 위한 재귀 함수
func validate(node *isTreeValidNode, low, high int) bool {
	if node == nil {
		return true
	}
	if node.Val <= low || node.Val >= high {
		return false
	}
	return validate(node.Left, low, node.Val) && validate(node.Right, node.Val, high)
}

func main() {
	// 유효한 이진 탐색 트리 예시
	root := &isTreeValidNode{Val: 2}
	root.Left = &isTreeValidNode{Val: 1}
	root.Right = &isTreeValidNode{Val: 3}

	fmt.Println(isValidBST(root)) // true

	// 유효하지 않은 이진 탐색 트리 예시
	root2 := &isTreeValidNode{Val: 5}
	root2.Left = &isTreeValidNode{Val: 1}
	root2.Right = &isTreeValidNode{Val: 4}
	root2.Right.Left = &isTreeValidNode{Val: 3}
	root2.Right.Right = &isTreeValidNode{Val: 6}

	fmt.Println(isValidBST(root2)) // false
}
