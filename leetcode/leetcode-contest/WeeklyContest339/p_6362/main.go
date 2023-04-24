package main

func main() {

}

func findTheLongestBalancedSubstring(s string) int {
	zero, one := 0, 0
	ans := 0

	for left, right, sz := 0, 0, len(s); right < sz; {
		for ; right < sz && s[right] == '0'; right++ {
			zero++
		}
		for ; right < sz && s[right] == '1' && zero != one; right++ {
			one++
		}
		// zero == one
		if zero == one {
			ans = max(ans, zero+one)
			left = right + 1
		} else {
			left++
		}

		right = left
		zero, one = 0, 0
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
