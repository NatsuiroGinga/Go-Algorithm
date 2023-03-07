package search

import "fmt"

func cover() {
	var n, m, r int
	_, _ = fmt.Scan(&n, &m, &r)
	mp := make([][]bool, n+5)
	for i := range mp {
		mp[i] = make([]bool, n+5)
	}

	for i := 0; i < m; i++ {
		var x, y int
		_, _ = fmt.Scan(&x, &y)

		for j := 1; j <= n; j++ {
			for k := 1; k <= n; k++ {
				mp[j][k] = true
			}
		}
	}

	count := 0
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if mp[i][j] {
				count++
			}
		}
	}

	fmt.Println(count)
}
