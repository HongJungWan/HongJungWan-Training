// TODO: 이진 탐색 트리
// TODO: 유효한 이진 탐색 트리 확인 (DFS)

package main

import (
	"fmt"
	"math"
)

type isTreeValid struct {
	Val   int
	Left  *isTreeValid
	Right *isTreeValid
}

// 유효한 이진 탐색 트리인지 확인하는 함수
func isValidBST(root *isTreeValid) bool {
	return validateDFS(root, math.MinInt64, math.MaxInt64)
}

// 유효성 검사를 위한 재귀 함수
func validateDFS(node *isTreeValid, low, high int) bool {
	if node == nil {
		return true
	}
	if node.Val <= low || node.Val >= high {
		return false
	}
	return validateDFS(node.Left, low, node.Val) && validateDFS(node.Right, node.Val, high)
}

func main() {
	// 유효한 이진 탐색 트리 예시
	root := &isTreeValid{Val: 2}
	root.Left = &isTreeValid{Val: 1}
	root.Right = &isTreeValid{Val: 3}

	fmt.Println(isValidBST(root)) // true

	// 유효하지 않은 이진 탐색 트리 예시
	root2 := &isTreeValid{Val: 5}
	root2.Left = &isTreeValid{Val: 1}  // 왼쪽 자식 1: 1 < 5 (유효)
	root2.Right = &isTreeValid{Val: 4} // 오른쪽 자식 4: 4 > 5 (유효하지 않음)

	root2.Right.Left = &isTreeValid{Val: 3}
	root2.Right.Right = &isTreeValid{Val: 6}

	fmt.Println(isValidBST(root2)) // false
}

/*
유효한 이진 탐색 트리 규칙 : 모든 왼쪽 자식 노드는 부모보다 작고, 모든 오른쪽 자식 노드는 부모보다 큽니다.
*/
