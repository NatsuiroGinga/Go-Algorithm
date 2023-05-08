package main

func findLength(nums1, nums2 []int) int {
	var dp [1000 + 5]int
	ans := 0

	for i := 1; i <= len(nums1); i++ {
		for j := len(nums2); j >= 1; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
				if dp[j] > ans {
					ans = dp[j]
				}
			} else {
				dp[j] = 0
			}
		}
	}

	return ans
}

func main() {

}
