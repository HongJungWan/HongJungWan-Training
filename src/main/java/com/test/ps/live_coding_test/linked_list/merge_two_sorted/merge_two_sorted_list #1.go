// TODO: 두개의 연결 리스트 병합
// TODO: 단일 연결 리스트 (singly linked list)

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	listOne := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	listTwo := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	mergedList := mergeTwoLists(listOne, listTwo)
	for mergedList != nil {
		fmt.Print(mergedList.Val, " ")
		mergedList = mergedList.Next
	}
}

func mergeTwoLists(listOne *ListNode, listTwo *ListNode) *ListNode {
	dummy := &ListNode{} // 더미 노드는 실제 데이터가 없는 노드로, 병합된 리스트의 시작 부분을 가리키기 위한 목적으로 사용
	current := dummy     // 병합된 리스트의 현재 위치를 가리키는 포인터

	for listOne != nil && listTwo != nil {
		if listOne.Val < listTwo.Val {
			current.Next = listOne // current.Next에 listOne의 현재 노드를 저장
			listOne = listOne.Next // listOne 포인터를 한 칸 앞으로 이동
		} else {
			current.Next = listTwo // current.Next에 listTwo의 현재 노드를 저장
			listTwo = listTwo.Next // listTwo 포인터를 한 칸 앞으로 이동
		}
		current = current.Next // 포인터를 앞으로 이동시켜 다음 노드를 추가할 위치를 항상 가리키는 역할
	}

	// 남아있는 노드가 있다면 연결
	if listOne != nil {
		current.Next = listOne // 남아있는 listOne의 노드를 연결
	} else {
		current.Next = listTwo // 남아있는 listTwo의 노드를 연결
	}

	// 더미 노드의 다음 노드가 병합된 리스트의 헤드
	return dummy.Next
}

/*
[시작 상태]
listOne: 1 -> 2 -> 4
listTwo: 1 -> 3 -> 4
dummy: 빈 노드 (초기화 상태)
current: dummy



[첫 번째 반복]
listOne.Val (1) vs listTwo.Val (1)

listOne: 1 -> 2 -> 4
listTwo: 3 -> 4
dummy: 0 -> 1
current: 1



[두 번째 반복]
listOne.Val (1) vs listTwo.Val (3)

listOne: 2 -> 4
listTwo: 3 -> 4
dummy: 0 -> 1 -> 1
current: 1



[세 번째 반복]
listOne.Val (2) vs listTwo.Val (3)

listOne: 4
listTwo: 3 -> 4
dummy: 0 -> 1 -> 1 -> 2
current: 2



[네 번째 반복]
listOne.Val (4) vs listTwo.Val (3)

listOne: 4
listTwo: 4
dummy: 0 -> 1 -> 1 -> 2 -> 3
current: 3



[다섯 번째 반복]
listOne.Val (4) vs listTwo.Val (4)

listOne: 4
listTwo: nil (빈 리스트)
dummy: 0 -> 1 -> 1 -> 2 -> 3 -> 4
current: 4



[루프 종료 후]
listTwo가 빈 리스트이므로, listOne의 남은 노드를 current.Next에 연결

listOne: 4
listTwo: nil (빈 리스트)
dummy: 0 -> 1 -> 1 -> 2 -> 3 -> 4 -> 4
current: 4



최종 병합된 리스트는 dummy.Next에서 시작, 결과는 1 -> 1 -> 2 -> 3 -> 4 -> 4
*/
