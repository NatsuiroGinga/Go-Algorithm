package main

/*var tmp []int
var ans [][]int
var n int
var vis map[int]struct{}

func permute(nums []int) [][]int {
	ans = make([][]int, 0)
	n = len(nums)
	tmp = make([]int, 0)
	vis = make(map[int]struct{})

	back(nums)
	return ans
}

func back(nums []int) {
	if len(tmp) == n {
		nArr := make([]int, n)
		copy(nArr, tmp)
		ans = append(ans, nArr)
		return
	}
	for i := 0; i < n; i++ {
		if _, ok := vis[i]; ok {
			continue
		}

		vis[i] = struct{}{}
		tmp = append(tmp, nums[i])
		back(nums)
		tmp = tmp[:len(tmp)-1]
		delete(vis, i)
	}
}
*/
