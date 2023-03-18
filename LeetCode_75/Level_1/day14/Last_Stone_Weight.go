package day14

import (
	"container/heap"
	"fmt"
)

func lastStoneWeight(stones []int) int {
	h := &IntHeap{}
	heap.Init(h)

	for _, stone := range stones {
		heap.Push(h, stone)
	}

	for h.Len() > 1 {
		y := heap.Pop(h)
		x := heap.Pop(h)

		fmt.Println("y:", y)
		fmt.Println("x:", x)

		if y != x {
			heap.Push(h, y.(int)-x.(int))
		}
	}

	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	l := len(*h)
	x := (*h)[l-1]
	*h = (*h)[:l-1]
	return x
}
