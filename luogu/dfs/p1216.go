package main

import (
	"bufio"
	"fmt"
	"os"
)

var input = bufio.NewReader(os.Stdin)
var output = bufio.NewWriter(os.Stdout)

const MAXN = 20

var n int
var ans [MAXN * 2]int
var ok [3][MAXN * 2]bool
var count int

func main() {
	defer output.Flush()
	fmt.Fscan(input, &n)
	dfs(1)
	fmt.Fprintln(output, count)
}

func dfs(cur int) {
	if cur > n {
		count++
		if count <= 3 {
			pr()
		}
		return
	}
	for i := 1; i <= n; i++ {
		if !ok[0][i] && !ok[1][i+cur] && !ok[2][i+n-cur] {
			ok[0][i] = true
			ok[1][i+cur] = true
			ok[2][i+n-cur] = true
			ans[cur] = i
			dfs(cur + 1)
			ok[0][i] = false
			ok[1][i+cur] = false
			ok[2][i+n-cur] = false
		}
	}
}

func pr() {
	for i := 1; i <= n; i++ {
		fmt.Fprintf(output, "%d ", ans[i])
	}
	fmt.Fprintln(output)
}
