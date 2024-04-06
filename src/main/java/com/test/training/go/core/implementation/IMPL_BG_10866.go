package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Deque struct 정의
type Deque struct {
	items []int
}

func (d *Deque) push_front(item int) {
	d.items = append([]int{item}, d.items...)
}

func (d *Deque) push_back(item int) {
	d.items = append(d.items, item)
}

func (d *Deque) pop_front() int {
	if len(d.items) == 0 {
		return -1
	}
	front := d.items[0]
	d.items = d.items[1:]
	return front
}

func (d *Deque) pop_back() int {
	if len(d.items) == 0 {
		return -1
	}
	back := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return back
}

func (d *Deque) size() int {
	return len(d.items)
}

func (d *Deque) empty() int {
	if len(d.items) == 0 {
		return 1
	}
	return 0
}

func (d *Deque) front() int {
	if len(d.items) == 0 {
		return -1
	}
	return d.items[0]
}

func (d *Deque) back() int {
	if len(d.items) == 0 {
		return -1
	}
	return d.items[len(d.items)-1]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	fmt.Scanln(&n)

	deque := Deque{}

	for i := 0; i < n; i++ {
		scanner.Scan()
		command := scanner.Text()
		parts := strings.Fields(command)

		switch parts[0] {
		case "push_front":
			num, _ := strconv.Atoi(parts[1])
			deque.push_front(num)
		case "push_back":
			num, _ := strconv.Atoi(parts[1])
			deque.push_back(num)
		case "pop_front":
			fmt.Println(deque.pop_front())
		case "pop_back":
			fmt.Println(deque.pop_back())
		case "size":
			fmt.Println(deque.size())
		case "empty":
			fmt.Println(deque.empty())
		case "front":
			fmt.Println(deque.front())
		case "back":
			fmt.Println(deque.back())
		}
	}
}
