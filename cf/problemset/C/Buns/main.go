package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, m, c, d, a, b int
	dp               []int
)

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	fmt.Fscan(gin, &n, &m, &c, &d)
	dp = make([]int, n+5)
	for i := 1; i <= n/c; i++ {
		dp[i*c] = i * d
	}

	for ; m > 0; m-- {
		fmt.Fscan(gin, &a, &b, &c, &d)
		for i := 1; i <= a/b; i++ {
			for j := n; j >= c; j-- {
				dp[j] = max(dp[j], dp[j-c]+d)
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans = max(ans, dp[i])
	}

	fmt.Fprintln(gout, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
