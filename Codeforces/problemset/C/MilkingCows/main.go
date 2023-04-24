package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, ans int
	cows   []bool
)

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	fmt.Fscan(gin, &n)
	cows = make([]bool, n)
	count := 0

	for i := range cows { // 0 0 1 0, 1 0 1 0 1
		fmt.Fscan(gin, &cows[i])

		if cows[i] {
			count++
		} else {
			ans += count
		}
	}

	fmt.Fprintln(gout, ans)
}
