package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n, m, k, x, y, a, b int
	fmt.Fscan(in, &t, &n, &m, &k, &x, &y)
	s := (x + y) % 2

	for i := 0; i < k; i++ {
		fmt.Fscan(in, &a, &b)
		if (a+b)%2 == s {
			fmt.Fprintln(out, "YES")
			return
		}
	}

	fmt.Fprintln(out, "NO")
}
