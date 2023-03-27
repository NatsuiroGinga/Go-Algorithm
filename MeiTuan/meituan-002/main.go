package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n      int
	w      []int
	preSum []int
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	w = make([]int, n+5)
	preSum = make([]int, n+5)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &w[i])
		preSum[i] = preSum[i-1] + w[i]
	}

	for i := 1; i <= n; i++ {
		var take, ans int
		fmt.Fscan(in, &take)
		ans = Max(0, preSum[take-1])

		for j := take + 1; j <= n; j++ {
			preSum[j] -= preSum[take]
			ans = Max(preSum[j], ans)
		}

		preSum[take] = 0
		fmt.Fprintln(out, ans)
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
