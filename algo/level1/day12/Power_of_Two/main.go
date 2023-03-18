package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(isPowerOfTwo(n))
}

func isPowerOfTwo(n int) bool {
	formatInt := strconv.FormatInt(int64(n), 2)
	return n > 0 && strings.Count(formatInt, "1") == 1
}
