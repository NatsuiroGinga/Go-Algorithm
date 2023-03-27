package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	t, m, a, b uint64
	dp         []uint64
)

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	fmt.Fscan(gin, &t, &m)
	dp = make([]uint64, t+5)

	var i uint64
	for i = 1; i <= m; i++ {
		fmt.Fscan(gin, &a, &b)
		for j := a; j <= t; j++ {
			dp[j] = max(dp[j], dp[j-a]+b)
		}
	}

	fmt.Fprintln(gout, dp[t])
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
