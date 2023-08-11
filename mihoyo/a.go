package main

import (
	"fmt"
)

var t, n, ans int

func main() {
	fmt.Scan(&t)
	for ; t > 0; t-- {
		solve()
	}
	fmt.Println(ans)
}

func solve() {
	fmt.Scan(&n)
	var q []string
	friend := 0
	ok := true
	var s string

	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		q = append(q, s)
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		if s == q[i] {
			friend++
		} else {
			friend--
			if friend < 0 {
				ok = false
			}
		}
	}

	if ok {
		ans++
	}
}
