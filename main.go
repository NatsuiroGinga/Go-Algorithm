package main

import "fmt"

func main() {
	var nodes []Node
	nodes = append(nodes, Node{1, 2, 3})

	for i := range nodes {
		fmt.Println(nodes[i])
	}

	fmt.Println(nodes)
}

type Node struct {
	x, y, step int
}
