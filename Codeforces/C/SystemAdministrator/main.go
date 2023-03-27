package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, m, v int
)

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	fmt.Fscan(gin, &n, &m, &v)
	if m < n-1 || m > (n-1)*(n-2)/2+1 {
		fmt.Fprintln(gout, -1)
		return
	}

	m -= n - 1
	for i := 1; i <= n; i++ {
		if i != v {
			fmt.Fprintln(gout, v, i)
		}
	}

	if m > 0 {
		for i := 2; i < n; i++ {
			if i == v {
				continue
			}
			for j := 1; j < i; j++ {
				if j == v {
					continue
				}
				fmt.Fprintln(gout, i, j)
				m--
				if m == 0 {
					return
				}
			}
		}
	}
}
