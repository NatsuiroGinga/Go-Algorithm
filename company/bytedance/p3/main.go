package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var s string

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	fmt.Fscan(gin, &s) // (x-1)(x+5)

	split := strings.Split(s, ")")
	split = split[:len(split)-1]
	pre := int64(1)
	var pos int64

	for i, sub := range split {
		parseInt, _ := strconv.ParseInt(sub[2:], 10, 64)
		m := parseInt

		if i == len(split)-1 {
			pos = m
			break
		}

		pre = ((pre % 10007) * (m % 10007)) % 10007
	}

	ans := (pos + pre) % 10007

	fmt.Fprintln(gout, ans)
}
