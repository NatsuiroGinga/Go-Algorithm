package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	nums      [4]int64
	operators [3]string
	ans       int64
	vis       [4]bool
)

func main() {
	ans = math.MaxInt64
	gin := bufio.NewReader(os.Stdin)
	for i := 0; i < 4; i++ {
		fmt.Scan(&nums[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Fscan(gin, &operators[i])
	}
	dfs(0)

	fmt.Println(ans)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func dfs(cur int) {
	if cur == 3 {
		for i := 0; i < 4; i++ {
			if !vis[i] {
				ans = min(ans, nums[i])
			}
		}
		return
	}

	for i := 0; i < 4; i++ {
		if vis[i] {
			continue
		}

		for j := i + 1; j < 4; j++ {
			if vis[j] {
				continue
			}

			tmp := nums[i]
			vis[j] = true

			if operators[cur] == "+" {
				nums[i] += nums[j]
			} else {
				nums[i] *= nums[j]
			}

			if nums[i] < ans {
				dfs(cur + 1)
			}

			vis[j] = false
			nums[i] = tmp
		}
	}
}
