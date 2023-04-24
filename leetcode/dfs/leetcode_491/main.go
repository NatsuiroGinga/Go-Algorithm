package main

import (
	"fmt"
)

var (
	ans  [][]int
	path []int
)

func findSubsequences(nums []int) [][]int {
	ans = make([][]int, 0)
	path = make([]int, 0)
	dfs(0, nums)

	return ans
}

func dfs(cur int, nums []int) {
	if len(path) >= 2 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)
	}

	set := map[int]struct{}{}

	for i := cur; i < len(nums); i++ {
		// 保证递增
		if len(path) > 0 && path[len(path)-1] > nums[i] {
			continue
		}
		// 去重
		if _, ok := set[nums[i]]; ok {
			continue
		}

		set[nums[i]] = struct{}{}
		path = append(path, nums[i])
		dfs(i+1, nums)
		path = path[:len(path)-1]
	}
}

func main() {
	nums := []int{4, 6, 7, 7}
	ans := findSubsequences(nums)
	fmt.Println(ans)
}
