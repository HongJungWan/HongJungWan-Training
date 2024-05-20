// TODO: 스택 2개로 큐 만들기
// TODO:

package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{
		items: []int{},
	}
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	lastIndex := len(s.items) - 1
	return s.items[lastIndex], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Queue 구조체 정의
type Queue struct {
	stack1 *Stack
	stack2 *Stack
}

// NewQueue 큐를 초기화하는 함수
func NewQueue() *Queue {
	return &Queue{
		stack1: NewStack(),
		stack2: NewStack(),
	}
}

// Enqueue 큐에 요소를 추가하는 함수
func (q *Queue) Enqueue(item int) {
	q.stack1.Push(item)
}

// Dequeue 큐에서 요소를 제거하고 반환하는 함수
func (q *Queue) Dequeue() (int, error) {
	if q.stack2.IsEmpty() {
		for !q.stack1.IsEmpty() {
			item, _ := q.stack1.Pop()
			q.stack2.Push(item)
		}
	}
	return q.stack2.Pop()
}

// Peek 큐의 첫 번째 요소를 반환하는 함수
func (q *Queue) Peek() (int, error) {
	if q.stack2.IsEmpty() {
		for !q.stack1.IsEmpty() {
			item, _ := q.stack1.Pop()
			q.stack2.Push(item)
		}
	}
	return q.stack2.Peek()
}

// IsEmpty 큐가 비어있는지 확인하는 함수
func (q *Queue) IsEmpty() bool {
	return q.stack1.IsEmpty() && q.stack2.IsEmpty()
}

func main() {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// 큐의 첫 번째 요소 출력
	front, _ := queue.Peek()
	fmt.Println("Front element:", front) // 출력: Front element: 1

	// 큐에서 요소 제거 및 출력
	item, _ := queue.Dequeue()
	fmt.Println("Dequeued element:", item) // 출력: Dequeued element: 1

	// 큐가 비어있는지 확인
	fmt.Println("Is queue empty?", queue.IsEmpty()) // 출력: Is queue empty? false
}

/* 아이디어

1. Stack을 만든다.
2. 포인터 타입 Stack 2개로 Queue 구조체를 만든다.
3. Stack 2개로 Element를 옮기면서 Queue 기능을 따라 만드는 것


[초기 상태]
Queue:
    stack1: []
    stack2: []


[Enqueue 1]
Queue:
    stack1: [1]
    stack2: []

[Enqueue 2]
Queue:
    stack1: [1, 2]
    stack2: []

[Enqueue 3]
Queue:
    stack1: [1, 2, 3]
    stack2: []


[Dequeue], `stack2`에서 pop
Queue:
    stack1: []
    stack2: [3, 2, 1]

Queue:
    stack1: []
    stack2: [3, 2]


[Peek], `stack2`의 top 요소 반환: `1`
Queue:
    stack1: []
    stack2: [3, 2, 1]
*/
