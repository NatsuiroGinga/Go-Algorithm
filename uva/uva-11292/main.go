package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	n, m          int
	heads, powers []int
)

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	for count, _ := fmt.Fscan(gin, &n, &m); count == 2; count, _ = fmt.Fscan(gin, &n, &m) {
		heads = make([]int, n)
		powers = make([]int, m)

		for i := 0; i < n; i++ {
			fmt.Fscan(gin, &heads[i])
		}

		for i := 0; i < m; i++ {
			fmt.Fscan(gin, &powers[i])
		}

		sort.Ints(heads)
		sort.Ints(powers)

		cur, minCost := 0, 0
		for _, knight := range powers {
			if knight >= heads[cur] {
				minCost += knight
				cur++
				if cur == n {
					break
				}
			}
		}

		if cur < n {
			fmt.Fprintln(gout, "Loowater is doomed!")
			return
		}

		fmt.Fprintln(gout, minCost)
	}
}
