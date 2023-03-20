package main

import "sort"

/*
1. Define a struct to store the original position and value of the array
2. Sort the array
3. Use sort.Search to find the target
4. If found, return the original position, else return -1
*/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	nodes := make(Nodes, len(nums))
	for i, v := range nums {
		nodes[i] = Node{pos: i, val: v}
	}
	sort.Sort(nodes)
	i := sort.Search(len(nodes), func(i int) bool {
		return nodes[i].val >= target
	})

	if i < len(nodes) && nodes[i].val == target {
		return nodes[i].pos
	}

	return -1
}

type Node struct {
	pos, val int
}

type Nodes []Node

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Less(i, j int) bool {
	return n[i].val < n[j].val
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
