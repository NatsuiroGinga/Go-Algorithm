package dfs

var (
	ans  [][]int
	path []int
)

func combinationSum3(k, n int) [][]int {
	ans = make([][]int, 0)
	path = make([]int, 0)
	dfs3(n, k, 1)
	return ans
}

func dfs3(n, k, cur int) {
	if n < 0 {
		return
	}
	if len(path) == k && n == 0 {
		subset := make([]int, k)
		copy(subset, path)
		ans = append(ans, subset)
		return
	}

	for i := cur; i <= 9-(k-len(path))+1; i++ {
		path = append(path, i)
		dfs3(n-i, k, i+1)
		path = path[:len(path)-1]
	}
}
