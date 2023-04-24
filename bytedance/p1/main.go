package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	var n int
	fmt.Fscan(gin, &n)

	preorder := make([]int, n)
	inorder := make([]int, n)

	for i := range preorder {
		fmt.Fscan(gin, &preorder[i])
	}

	for i := range inorder {
		fmt.Fscan(gin, &inorder[i])
	}

}
