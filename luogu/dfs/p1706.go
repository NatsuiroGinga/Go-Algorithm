package main

import (
	"bufio"
	"fmt"
	"os"
)

var input = bufio.NewReader(os.Stdin)
var output = bufio.NewWriter(os.Stdout)
var n int
var ans [10]int
var vis [10]bool

func main() {
	defer output.Flush()
	fmt.Fscan(input, &n)
	dfs(1)
}

func dfs(cur int) {
	if cur > n {
		pr()
		return
	}
	for i := 1; i <= n; i++ {
		if !vis[i] {
			vis[i] = true
			ans[cur] = i
			dfs(cur + 1)
			vis[i] = false
		}
	}
}

func pr() {
	for i := 1; i <= n; i++ {
		fmt.Fprintf(output, "%5d", ans[i])
	}
	fmt.Fprintln(output)
}
