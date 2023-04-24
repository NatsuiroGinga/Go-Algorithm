package main

import "fmt"

var (
	ans  [][]int
	path []int
	set  map[int]struct{}
)

func permute(nums []int) [][]int {
	ans = make([][]int, 0)
	path = make([]int, 0)
	set = map[int]struct{}{}
	dfs(nums)
	return ans
}

func dfs(nums []int) {
	if len(path) == len(nums) {
		t := make([]int, len(path))
		copy(t, path)
		ans = append(ans, t)
		return
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := set[nums[i]]; ok {
			continue
		}
		set[nums[i]] = struct{}{}
		path = append(path, nums[i])
		dfs(nums)
		path = path[:len(path)-1]
		delete(set, nums[i])
	}
}

func main() {
	nums := []int{1, 2, 3}
	ret := permute(nums)
	fmt.Println(ret)
}
