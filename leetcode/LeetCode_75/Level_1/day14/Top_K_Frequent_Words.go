package day14

import (
	"container/heap"
)

func topKFrequent(words []string, k int) []string {
	h := &WordHeap{}
	var ans []string
	mp := make(map[string]int, len(words))

	for _, word := range words {
		mp[word]++
	}

	for key, value := range mp {
		heap.Push(h, Word{key, value})
	}

	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(h).(Word).content)
	}

	return ans
}

type Word struct {
	content string
	fre     int
}

type WordHeap []Word

func (w *WordHeap) Len() int {
	return len(*w)
}

func (w *WordHeap) Less(i, j int) bool {
	if (*w)[i].fre == (*w)[j].fre {
		return (*w)[i].content < (*w)[j].content
	}
	return (*w)[i].fre > (*w)[j].fre
}

func (w *WordHeap) Swap(i, j int) {
	(*w)[i], (*w)[j] = (*w)[j], (*w)[i]
}

func (w *WordHeap) Push(x interface{}) {
	*w = append(*w, x.(Word))
}

func (w *WordHeap) Pop() interface{} {
	l := len(*w)
	x := (*w)[l-1]
	*w = (*w)[:l-1]
	return x
}
