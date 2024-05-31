// TODO: Floyd’s Cycle Detection Algorithm, 연결 리스트가 주어졌을 때 해당 연결 리스트가 nil로 끝나는지 순환 되는지 검사

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}

	// Connect nodes to form the list: 3 -> 2 -> 0 -> -4
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	// Create a cycle: -4 -> 2 (node4.Next = node2)
	node4.Next = node2

	// Test the hasCycle function
	fmt.Println(hasCycle(node1)) // Output: true

	// Create a list without a cycle: 1 -> 2
	nodeA := &ListNode{Val: 1}
	nodeB := &ListNode{Val: 2}
	nodeA.Next = nodeB

	// Test the hasCycle function
	fmt.Println(hasCycle(nodeA)) // Output: false
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

/*
Cycle이 있는 연결 리스트: 3 -> 2 -> 0 -> -4 -> (cycle back to 2)

Cycle이 없는 연결 리스트: 1 -> 2
*/
