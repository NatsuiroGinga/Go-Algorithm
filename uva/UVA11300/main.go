package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	var (
		n    int64
		a, c []int64
	)

	for {
		if _, err := fmt.Fscan(gin, &n); err != nil {
			break
		}

		a = make([]int64, n+5)
		c = make([]int64, n)

		var avg, i, sum int64
		for i = 1; i <= n; i++ {
			fmt.Fscan(gin, &a[i])
			sum += a[i]
		}
		avg = sum / n

		c[0] = 0
		for i = 1; i < n; i++ {
			c[i] = c[i-1] + a[i] - avg
		}

		sort.Slice(c, func(i, j int) bool {
			return c[i] < c[j]
		})
		mid := c[n/2]

		var ans int64
		for i = 0; i < n; i++ {
			ans += abs(c[i] - mid)
		}

		fmt.Fprintln(gout, ans)
	}
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
