package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	fmt.Fscan(gin, &n)
}
