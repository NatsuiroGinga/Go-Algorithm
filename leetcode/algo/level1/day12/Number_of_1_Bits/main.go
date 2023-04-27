package main

import (
	"fmt"
)

func main() {
	var n uint32
	fmt.Scan(&n)
	fmt.Println(hammingWeight(n))
}

func hammingWeight(num uint32) int {
	var count int
	for ; num > 0; count++ {
		num &= num - 1
	}
	return count
}
