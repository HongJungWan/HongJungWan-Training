// TODO: 연결 리스트
// TODO: 단일 연결 리스트 (singly linked list)

package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

/*
1. Head 필드는 연결 리스트의 첫 번째 노드를 가리킨다
2. 리스트가 비어 있으면 Head는 nil
3. 리스트에 새로운 노드를 추가하거나 탐색할 때, Head를 통해 리스트의 시작점에 접근
*/
type LinkedList struct {
	Head *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil}
}

func main() {
	linkedList := NewLinkedList()
	linkedList.AddToEnd(1)
	linkedList.AddToEnd(2)
	linkedList.AddToEnd(3)
	linkedList.AddToEnd(4)

	fmt.Println("The elements of the linked list are:")
	linkedList.PrintList() // Output: 1 2 3 4
}

func (linkedList *LinkedList) AddToEnd(value int) {
	newNode := &ListNode{Value: value}
	if linkedList.Head == nil {
		linkedList.Head = newNode
	} else {
		current := linkedList.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (linkedList *LinkedList) PrintList() {
	current := linkedList.Head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}
