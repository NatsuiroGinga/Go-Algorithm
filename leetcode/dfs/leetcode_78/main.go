package main

import "fmt"

var (
	ans  [][]int
	path []int
)

func subsets(nums []int) [][]int {
	ans = make([][]int, 0)
	path = make([]int, 0)
	dfs(0, nums)
	return ans
}

func dfs(cur int, nums []int) {
	subset := make([]int, len(path))
	copy(subset, path)
	ans = append(ans, subset)

	for i := cur; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs(i+1, nums)
		path = path[:len(path)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	ret := subsets(nums)
	fmt.Println(ret)
}
