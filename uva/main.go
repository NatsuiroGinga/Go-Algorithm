package main

import (
	"fmt"
	"sort"
)

var vis [1000 + 5]bool
var pri [1000 + 5]int
var cnt int

func prime(n int) {
	for i := 2; i <= n; i++ {
		if !vis[i] {
			pri[cnt] = i
			cnt++
		}
		for j := 0; j < cnt && i*pri[j] <= n; j++ {
			vis[i*pri[j]] = true
			if i%pri[j] == 0 {
				break
			}
		}
	}
}

func main() {
	a := []int{2, 2}
	fmt.Println(sort.IntsAreSorted(a))
}
