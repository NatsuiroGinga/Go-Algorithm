package main

import (
	"fmt"
	"sort"
)

func main() {
	it := [][]int{{1, 4}, {2, 3}}
	fmt.Println(merge(it))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{{intervals[0][0], intervals[0][1]}}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= ans[len(ans)-1][1] {
			ans[len(ans)-1][1] = max(intervals[i][1], ans[len(ans)-1][1])
		} else {
			ans = append(ans, []int{intervals[i][0], intervals[i][1]})
		}
	}

	return ans
}

var max = func(a, b int) int {
	if a > b {
		return a
	}
	return b
}
