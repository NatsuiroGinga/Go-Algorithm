package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, m int
	w, d int
	dp   []int
)

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	fmt.Fscan(gin, &n, &m)
	dp = make([]int, m+5)

	for i := 1; i <= n; i++ {
		fmt.Fscan(gin, &w, &d)
		for j := m; j >= w; j-- {
			dp[j] = max(dp[j], dp[j-w]+d)
		}
	}

	fmt.Fprintln(gout, dp[m])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
