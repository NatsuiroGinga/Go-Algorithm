package main

import (
	"fmt"
)

var (
	ret  [][]int
	path []int
)

func combinationSum(candidates []int, target int) [][]int {
	ret = make([][]int, 0)
	path = make([]int, 0)
	dfs(target, 0, candidates)
	return ret
}

func dfs(targetSum, cur int, candidates []int) {
	if targetSum < 0 {
		return
	}
	if 0 == targetSum {
		subset := make([]int, len(path))
		copy(subset, path)
		ret = append(ret, subset)
		return
	}
	for i := cur; i < len(candidates); i++ {
		path = append(path, candidates[i])
		dfs(targetSum-candidates[i], i, candidates)
		path = path[:len(path)-1]
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
