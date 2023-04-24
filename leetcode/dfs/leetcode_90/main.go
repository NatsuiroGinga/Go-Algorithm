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

func subsetsWithDup(nums []int) [][]int {
	ans = make([][]int, 0)
	path = make([]int, 0)
	vis = make([]bool, len(nums))
	sort.Ints(nums)
	dfs(0, nums)

	return ans
}

func dfs(cur int, nums []int) {
	subset := make([]int, len(path))
	copy(subset, path)
	ans = append(ans, subset)

	for i := cur; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && !vis[i-1] {
			continue
		}

		path = append(path, nums[i])
		vis[i] = true
		dfs(i+1, nums)
		vis[i] = false
		path = path[:len(path)-1]
	}
}

func main() {
	nums := []int{1, 2, 2}
	dup := subsetsWithDup(nums)
	fmt.Println(dup)
}
