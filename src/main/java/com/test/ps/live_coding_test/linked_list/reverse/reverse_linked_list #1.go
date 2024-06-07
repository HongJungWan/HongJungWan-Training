// TODO: 연결 리스트 뒤집기
// TODO: 단일 연결 리스트 (singly linked list)

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Create a list: 1 -> 2 -> 3 -> nil
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}

	node1.Next = node2
	node2.Next = node3

	fmt.Println("Original list:")
	printList(node1)

	fmt.Println("\nReversed list:")
	printList(reverseList(node1))
}

func reverseList(head *ListNode) *ListNode {
	var previousNode *ListNode
	currentNode := head

	for currentNode != nil {
		next := currentNode.Next        // 현재 노드의 다음 노드를 저장
		currentNode.Next = previousNode // 현재 노드의 포인터를 이전 노드를 가리키도록 변경
		previousNode = currentNode      // 이전 노드를 현재 노드로 업데이트
		currentNode = next              // 현재 노드를 다음 노드로 이동
	}

	return previousNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

/*
초기 상태: previousNode = nil, currentNode = 1

[첫 번째 반복]
next = 2
1의 Next를 nil로 설정 (1 -> nil)
previousNode를 1로 이동 (previousNode = 1)
currentNode를 2로 이동 (currentNode = 2)


[두 번째 반복]
next = 3
2의 Next를 1로 설정 (2 -> 1 -> nil)
previousNode를 2로 이동 (previousNode = 2)
currentNode를 3로 이동 (currentNode = 3)


[세 번째 반복]
next = nil
3의 Next를 2로 설정 (3 -> 2 -> 1 -> nil)
previousNode를 3로 이동 (previousNode = 3)
currentNode를 nil로 이동 (currentNode = nil)
*/
