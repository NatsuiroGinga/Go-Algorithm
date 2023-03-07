package search

import (
	"fmt"
	"sort"
)

var (
	sticks = make([]int, 4) // store the length of sticks
	ans    = make([]int, 3) // store the length of the three sticks
)
var ret = "IMPOSSIBLE" // store the result

func main() {
	for i := 0; i < 4; i++ {
		_, _ = fmt.Scan(&sticks[i])
	}
	sort.Ints(sticks) // sort the sticks
	dfs(0, 0)
	fmt.Print(ret)
}

func dfs(cur int, pos int) { // cur: count of chosen sticks, pos: the start position of the next stick
	if cur == 3 { // find a combination
		if ans[0]+ans[1] > ans[2] { // check if it is a triangle
			ret = "TRIANGLE"
		} else if ans[0]+ans[1] == ans[2] && ret != "TRIANGLE" {
			/*
				check if it is a segment and the result is not TRIANGLE
					 if the result is TRIANGLE, it should not be changed
			*/
			ret = "SEGMENT"
		}
		return
	}

	for i := pos; i < 4; i++ {
		ans[cur] = sticks[i] // choose the i-th stick
		dfs(cur+1, i+1)      // search the next stick
	}
}
