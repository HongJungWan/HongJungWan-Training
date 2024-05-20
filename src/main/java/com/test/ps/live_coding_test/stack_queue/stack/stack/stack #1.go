// TODO: 스택 만들기
// TODO: Basic

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

func main() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// 스택의 마지막 요소 출력
	top, _ := stack.Peek()
	fmt.Println("Top element:", top) // 출력: Top element: 3

	// 스택에서 요소 제거 및 출력
	item, _ := stack.Pop()
	fmt.Println("Popped element:", item) // 출력: Popped element: 3

	// 스택이 비어있는지 확인
	fmt.Println("Is stack empty?", stack.IsEmpty()) // 출력: Is stack empty? false
}
