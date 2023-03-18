package main

/*var ans [][]int
var tmp []int

func combine(n int, k int) [][]int {
	ans = make([][]int, 0)
	tmp = make([]int, 0)

	dfs(1, n, k)
	return ans
}

func dfs(start, n, k int) {
	if k == 0 {
		nArr := make([]int, len(tmp))
		copy(nArr, tmp)
		ans = append(ans, nArr)
		return
	}
	for i := start; i <= n; i++ {
		tmp = append(tmp, i)
		dfs(i+1, n, k-1)
		tmp = (tmp)[:len(tmp)-1]
	}
}
*/
