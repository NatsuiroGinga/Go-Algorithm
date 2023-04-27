package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	t, l, r int
)

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	fmt.Fscan(gin, &t)

	for ; t > 0; t-- {
		fmt.Fscan(gin, &l, &r)
		var luckiest int
		cur := math.MinInt

		for i := l; i <= r; i++ {
			numStr := strconv.Itoa(i)
			if strings.ContainsRune(numStr, '9') && strings.ContainsRune(numStr, '0') {
				luckiest = i
				break
			} else {
				num := i
				maxDigit, minDigit := math.MinInt, math.MaxInt

				for ; num != 0; num /= 10 {
					unit := num % 10
					maxDigit = max(maxDigit, unit)
					minDigit = min(minDigit, unit)
				}

				if maxDigit-minDigit > cur {
					cur = maxDigit - minDigit
					luckiest = i
				}
			}
		}

		fmt.Fprintln(gout, luckiest)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
