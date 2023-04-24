package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	count := 0

	sort.Ints(g)
	sort.Ints(s)

	for _, v := range s {
		if v >= g[count] {
			count++
			if count == len(g) {
				break
			}
		}
	}

	return count
}

func main() {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	children := findContentChildren(g, s)
	fmt.Println(children)
}
