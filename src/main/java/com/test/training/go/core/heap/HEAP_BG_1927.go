package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

// IntHeap is a type that implements heap.Interface
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		if x == 0 {
			if h.Len() == 0 {
				writer.WriteString("0\n")
			} else {
				fmt.Fprintf(writer, "%d\n", (*h)[0])
				heap.Remove(h, 0)
			}
		} else {
			heap.Push(h, x)
		}
	}
}
