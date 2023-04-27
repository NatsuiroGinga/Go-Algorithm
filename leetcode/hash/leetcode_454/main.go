package leetcode_454

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	ans := 0
	n := len(nums1)
	sum := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum[nums1[i]+nums2[j]]++
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if v, ok := sum[0-nums3[i]-nums4[j]]; ok {
				ans += v
			}
		}
	}

	return ans
}
