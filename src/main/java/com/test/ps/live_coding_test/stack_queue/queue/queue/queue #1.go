// TODO: 큐 만들기
// TODO: Basic

package queue

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

func main() {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	front, _ := queue.Peek()
	fmt.Println("Front element:", front) // 출력: Front element: 1

	item, _ := queue.Dequeue()
	fmt.Println("Dequeued element:", item)          // 출력: Dequeued element: 1
	fmt.Println("Is queue empty?", queue.IsEmpty()) // 출력: Is queue empty? false
}
