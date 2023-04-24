package main

import (
	"fmt"
	"sort"
)

var (
	ret  [][]int
	path []int
	vis  []bool
)

func combinationSum2(candidates []int, target int) [][]int {
	ret = make([][]int, 0)
	path = make([]int, 0)
	vis = make([]bool, len(candidates))
	sort.Ints(candidates)
	dfs(target, 0, candidates)
	return ret
}

func dfs(targetSum, cur int, candidates []int) {
	if targetSum < 0 {
		return
	}
	if targetSum == 0 {
		subset := make([]int, len(path))
		copy(subset, path)
		ret = append(ret, subset)
		return
	}
	for i := cur; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] && !vis[i-1] {
			continue
		}
		path = append(path, candidates[i])
		vis[i] = true
		dfs(targetSum-candidates[i], i+1, candidates)
		vis[i] = false
		path = path[:len(path)-1]
	}
}

func main() {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	fmt.Println(combinationSum2(candidates, target))
}
