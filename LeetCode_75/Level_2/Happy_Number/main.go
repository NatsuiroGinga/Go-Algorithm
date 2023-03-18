package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(isHappy(n))
}

func isHappy(n int) bool {
	sum := ComputeSum(n)
	set := make(map[int]struct{})
	set[sum] = struct{}{}

	for sum != 1 {
		sum = ComputeSum(sum)

		if _, ok := set[sum]; ok {
			return false
		}

		set[sum] = struct{}{}
	}

	return true
}

func ComputeSum(n int) int {
	var sum int
	for sum = 0; n != 0; n /= 10 {
		x := n % 10
		sum += x * x
	}
	return sum
}
