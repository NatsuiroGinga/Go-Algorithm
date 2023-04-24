package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	t, n, m int
	cards   [][]int
)

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	fmt.Fscan(gin, &t)

	for ; t > 0; t-- {
		fmt.Fscan(gin, &n, &m)
		cards = make([][]int, n)

		for i := 0; i < n; i++ {
			cards[i] = make([]int, m)
			for j := 0; j < m; j++ {
				fmt.Fscan(gin, &cards[i][j])
			}
		}

		total := 0
		for i := 0; i < m; i++ {
			maxV, minV := cards[0][i], cards[0][i]
			for j := 0; j < n; j++ {
				maxV = max(maxV, cards[j][i])
				minV = min(minV, cards[j][i])
			}
			fmt.Fprintln(gout, "max =", maxV)
			fmt.Fprintln(gout, "min =", minV)
			total += 2 * (maxV - minV)
		}

		fmt.Fprintln(gout, total)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
