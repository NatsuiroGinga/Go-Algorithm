package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var (
	ts, te     string
	n, t, c, p int
	dp         []int
)

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	fmt.Fscan(gin, &ts, &te, &n)
	start, _ := time.Parse("15:04", ts)
	end, _ := time.Parse("15:04", te)
	minutes := int(end.Sub(start).Minutes())
	dp = make([]int, minutes+10)

	for i := 1; i <= n; i++ {
		fmt.Fscan(gin, &t, &c, &p)
		switch p {
		case 1:
			continue // 0-1背包
		case 0:
			continue // 完全背包
		default:
			continue // 多重背包
		}
	}

	fmt.Fprintln(gout, dp[minutes])
}
