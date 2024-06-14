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

/*
     5
    / \
   1   4
     /  \
	3    6
*/

/*
validateDFS(node *isTreeValid, low, high int) 호출:

1-1. [최초 validateDFS 호출] validateDFS(root2, math.MinInt64, math.MaxInt64) 호출
1-2. root2의 값은 5, 범위(low, high)는 (math.MinInt64, math.MaxInt64)


2-1. [왼쪽 자식 노드 기준, validateDFS 호출] validateDFS(1, math.MinInt64, 5) 호출
2-2. root2.Left의 값은 1, 범위는 (math.MinInt64, 5)
2-3. [node.Val <= low || node.Val >= high] 조건 검사, node.val == 1
2-4. root2.Left의 왼쪽과 오른쪽 자식 노드는 nil이므로 true 반환


3-1. [오른쪽 자식 노드 기준, validateDFS 호출] validateDFS(4, 5, math.MaxInt64) 호출
3-2. root2.Right의 값은 4, 범위는 (5, math.MaxInt64)
3-3. [node.Val <= low || node.Val >= high] 조건 검사, node.val == 4
3-4. 4는 유효하지 않은 값이므로 false를 반환
*/
