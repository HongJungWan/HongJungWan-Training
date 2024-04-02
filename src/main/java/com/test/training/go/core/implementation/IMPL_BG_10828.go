package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	elements []int
}

func (s *Stack) Push(x int) {
	s.elements = append(s.elements, x)
}

func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		return -1
	}
	lastIndex := len(s.elements) - 1
	topElement := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return topElement
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Empty() int {
	if len(s.elements) == 0 {
		return 1
	}
	return 0
}

func (s *Stack) Top() int {
	if len(s.elements) == 0 {
		return -1
	}
	return s.elements[len(s.elements)-1]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	fmt.Scan(&n)

	stack := Stack{}

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := scanner.Text()
		command := strings.Split(input, " ")

		switch command[0] {
		case "push":
			x, _ := strconv.Atoi(command[1])
			stack.Push(x)
		case "pop":
			fmt.Println(stack.Pop())
		case "size":
			fmt.Println(stack.Size())
		case "empty":
			fmt.Println(stack.Empty())
		case "top":
			fmt.Println(stack.Top())
		}
	}
}
