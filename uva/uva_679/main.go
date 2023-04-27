package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, ball, ans, t int

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	fmt.Fscan(gin, &t)
	for ; t > 0; t-- {
		fmt.Fscan(gin, &n, &ball)
		nodes := 1<<n - 1
		tree := make([]bool, nodes+1)

		for i := 1; i <= ball; i++ {
			j := 1

			for j <= nodes {
				tree[j] = !tree[j]

				if j*2 > nodes {
					break
				}

				j *= 2
				if tree[j] {
					j++
				}
			}

			ans = j
		}

		fmt.Fprintln(gout, ans)
	}
}
