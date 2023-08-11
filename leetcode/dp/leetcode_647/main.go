package leetcode_647

func countSubstrings(s string) (ans int) {
	dp := make([]bool, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[j] = true
				} else {
					dp[j] = dp[j-1]
				}

				if dp[j] {
					ans++
				}
			} else { // 必须要有这个else，否则dp[j]会被上一次的dp[j]覆盖
				dp[j] = false
			}
		}
	}

	return
}
