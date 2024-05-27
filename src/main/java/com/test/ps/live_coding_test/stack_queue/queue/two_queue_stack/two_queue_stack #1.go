// TODO: 큐 2개로 스택 만들기
// TODO: Push 메서드는 단순히 queue1에 요소를 추가
// TODO: Pop 메서드는 queue1의 요소를 queue2로 옮기면서 마지막 요소를 제거하여 반환
// TODO: Peek 메서드는 queue1의 요소를 queue2로 옮기면서 마지막 요소를 확인하고 다시 queue2로 옮긴다.

package two_queue_stack

import (
	"fmt"
)

type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{
		items: []int{},
	}
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

type Stack struct {
	queue1 *Queue
	queue2 *Queue
}

func NewStack() *Stack {
	return &Stack{
		queue1: NewQueue(),
		queue2: NewQueue(),
	}
}

func (s *Stack) Push(item int) {
	s.queue1.Enqueue(item)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}

	for len(s.queue1.items) > 1 {
		item, _ := s.queue1.Dequeue()
		s.queue2.Enqueue(item)
	}

	poppedItem, _ := s.queue1.Dequeue()
	s.queue1, s.queue2 = s.queue2, s.queue1

	return poppedItem, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}

	for len(s.queue1.items) > 1 {
		item, _ := s.queue1.Dequeue()
		s.queue2.Enqueue(item)
	}

	topItem, _ := s.queue1.Dequeue()
	s.queue2.Enqueue(topItem)
	s.queue1, s.queue2 = s.queue2, s.queue1

	return topItem, nil
}

func (s *Stack) IsEmpty() bool {
	return s.queue1.IsEmpty()
}

func main() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	top, _ := stack.Peek()
	fmt.Println("Top element:", top) // 출력: Top element: 3

	item, _ := stack.Pop()
	fmt.Println("Popped element:", item)            // 출력: Popped element: 3
	fmt.Println("Is stack empty?", stack.IsEmpty()) // 출력: Is stack empty? false
}
