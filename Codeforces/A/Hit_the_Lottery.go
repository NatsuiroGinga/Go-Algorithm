package main

import (
	"fmt"
)

func main() {
	var n int
	ans := 0
	fmt.Scan(&n)
	bills := []int{1, 5, 10, 20, 100}

	for i := len(bills) - 1; i >= 0; i-- {
		ans += n / bills[i]
		n = n % bills[i]
	}

	fmt.Println(ans)
}
