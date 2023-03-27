package main

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"os"
)

var (
	stack   list.List
	dp      []bool
	maxLen  int
	records map[int]int
)

func main() {
	in := bufio.NewReader(os.Stdin)
	records = make(map[int]int)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	seq, _ := in.ReadBytes('\n')
	n := len(seq)
	seq = bytes.TrimSpace(seq)
	dp = make([]bool, n+1)

	for i, b := range seq {
		if b == '(' {
			stack.PushBack(i)
		} else {
			if stack.Len() > 0 {
				dp[i] = true
				dp[stack.Back().Value.(int)] = true
				stack.Remove(stack.Back())
			}
		}
	}

	curLen := 0
	for _, b := range dp {
		if b {
			curLen++
		} else {
			if curLen >= maxLen {
				maxLen = curLen
				records[curLen]++
			}
			curLen = 0
		}
	}

	if maxLen > 0 {
		fmt.Fprintln(out, maxLen, records[maxLen])
	} else {
		fmt.Fprintln(out, 0, 1)
	}
}
