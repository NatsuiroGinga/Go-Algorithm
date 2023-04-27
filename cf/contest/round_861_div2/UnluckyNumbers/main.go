package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var t int
var l, r int64

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	fmt.Fscan(gin, &t)
	for ; t > 0; t-- {
		fmt.Fscan(gin, &l, &r)
		var luckiest int64
		var cur int64 = math.MaxInt

		for i := l; i <= r; i++ {
			numStr := strconv.FormatInt(i, 10)

			if len(numStr) == 1 || eq(numStr) {
				luckiest = i
				break
			} else {
				num := i
				var maxDigit, minDigit int64 = math.MinInt, math.MinInt

				for ; num != 0; num /= 10 {
					unit := num % 10
					maxDigit = max(maxDigit, unit)
				}

				num = i
				for ; num != 0; num /= 10 {
					var unit = num % 10
					if unit == maxDigit {
						continue
					}
					minDigit = max(minDigit, unit)
				}

				if maxDigit-minDigit < cur {
					cur = maxDigit - minDigit
					luckiest = i
				}
			}
		}

		fmt.Fprintln(gout, luckiest)
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func eq(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i+1] != s[i] {
			return false
		}
	}
	return true
}
