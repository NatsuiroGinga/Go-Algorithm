package dfs

func combine(n int, k int) [][]int {
	ans := &[][]int{}
	dfs(ans, 1, n)
	return *ans
}

func dfs(ans *[][]int, cur, n int) {
	if cur == n+1 {
		return
	}
	for i := cur + 1; i <= n; i++ {
		*ans = append(*ans, []int{cur, i})
		dfs(ans, i, n)
	}
}
