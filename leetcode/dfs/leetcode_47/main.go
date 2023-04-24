package main

import (
	"fmt"
	"sort"
)

var (
	ans  [][]int
	path []int
	vis  []bool
)

func permuteUnique(nums []int) [][]int {
	ans = make([][]int, 0)
	path = make([]int, 0)
	vis = make([]bool, len(nums))
	sort.Ints(nums)
	dfs(0, nums)
	return ans
}

func dfs(cur int, nums []int) {
	if len(path) == len(nums) {
		t := make([]int, len(path))
		copy(t, path)
		ans = append(ans, t)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 树层去重
		if i > 0 && nums[i] == nums[i-1] && !vis[i-1] {
			continue
		}
		// 树枝去重
		if vis[i] {
			continue
		}

		vis[i] = true
		path = append(path, nums[i])
		dfs(i+1, nums)
		path = path[:len(path)-1]
		vis[i] = false
	}
}

func main() {
	nums := []int{1, 2, 3}
	ret := permuteUnique(nums)
	fmt.Println(ret)
}
