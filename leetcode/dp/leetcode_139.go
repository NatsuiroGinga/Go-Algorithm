package dp

func wordBreak(s string, wordDict []string) bool {
	dp := [300 + 5]bool{true} // j个字符能否拼成
	set := map[string]struct{}{}

	for _, word := range wordDict {
		set[word] = struct{}{}
	}

	for j := 1; j <= len(s); j++ {
		for i := 0; i < j; i++ {
			if _, ok := set[s[i:j]]; ok && dp[i] {
				dp[j] = true
			}
		}
	}

	return dp[len(s)]
}
