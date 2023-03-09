package day9

func fib(n int) int {
	pre, cur, pos := 0, 1, 0

	if n == 1 {
		return cur
	}

	for i := 2; i <= n; i++ {
		pos = pre + cur
		pre = cur
		cur = pos
	}

	return pos
}
